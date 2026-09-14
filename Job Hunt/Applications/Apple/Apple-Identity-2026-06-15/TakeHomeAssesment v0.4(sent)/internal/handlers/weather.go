package handlers

import (
	"TakeHomeAssessment/internal/weather"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func GetWeatherHandler(locData <-chan weather.LocationInfo, timeOut time.Duration, locCircLog weather.CircularLog) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value("logger").(*slog.Logger)

		if !ok {
			logger = slog.Default()
		}

		ctx, cancelCtx := context.WithTimeout(r.Context(), timeOut)
		defer cancelCtx()

		for {
			select {
			case <-ctx.Done():

				// Degrade gracefully if the generator could not produce location data fast enough
				// attempt to use data from previous requests
				cachedLoc, ok := locCircLog.Read()

				if !ok {
					// All cache data is stale or there is none
					logger.Warn("Request timeout", slog.Int("HttpStatusCode", http.StatusRequestTimeout))
					http.Error(w, "Location data retrieval took too long.", http.StatusRequestTimeout)
					return
				}

				logger.Info("Request succesful", slog.Int("HttpStatusCode", http.StatusOK), slog.Bool("FromCache", true))
				w.WriteHeader(http.StatusOK)
				response := fmt.Sprintf("The weather in %v is:%v", cachedLoc.Name, cachedLoc.Forecast)
				w.Write([]byte(response))
				return

			case locInfo, ok := <-locData:
				if !ok {
					errString := "Location data not available."
					logger.Error(errString, slog.Int("HttpStatusCode", http.StatusInternalServerError))
					http.Error(w, "Location data not available.", http.StatusInternalServerError)
					return
				}

				if locInfo.ExpiresAt.Before(time.Now()) {
					// Data retrieved is stale, waiting for next fresh batch
					continue
				}

				logger.Info("Request succesful", slog.Int("HttpStatusCode", http.StatusOK), slog.Bool("FromCache", false))
				w.WriteHeader(http.StatusOK)
				response := fmt.Sprintf("The weather in %v is:%v", locInfo.Name, locInfo.Forecast)
				w.Write([]byte(response))
				return
			}
		}

	})
}

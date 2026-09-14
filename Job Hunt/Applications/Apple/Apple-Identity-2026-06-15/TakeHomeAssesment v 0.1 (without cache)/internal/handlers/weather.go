package handlers

import (
	"TakeHomeAssesment/internal/weather"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func GetWeatherHandler(locData <-chan weather.LocationInfo, timeOut time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value("logger").(*slog.Logger)

		if !ok {
			logger = slog.Default()
		}

		ctx, cancelCtx := context.WithTimeout(r.Context(), timeOut)
		defer cancelCtx()

		select {
		case <-ctx.Done():
			logger.Warn("Request timeout", slog.Int("HttpStatusCode", http.StatusRequestTimeout))
			http.Error(w, "Location data retrieval took too long.", http.StatusRequestTimeout)
			return
		case locInfo, ok := <-locData:
			if !ok {
				errString := "Location data not available."
				logger.Error(errString, slog.Int("HttpStatusCode", http.StatusInternalServerError))
				http.Error(w, "Location data not available.", http.StatusInternalServerError)
				return
			}

			logger.Info("Request succesful", slog.Int("HttpStatusCode", http.StatusOK))
			w.WriteHeader(http.StatusOK)
			response := fmt.Sprintf("The weather in %v is:%v", locInfo.Name, locInfo.Forecast)
			w.Write([]byte(response))
		}
	})
}

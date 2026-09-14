package main

import (
	"TakeHomeAssesment/internal/handlers"
	"TakeHomeAssesment/internal/middleware"
	"TakeHomeAssesment/internal/weather"
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	// Configure application logging
	logOptions := &slog.HandlerOptions{
		//Level: slog.LevelWarn,
		Level: slog.LevelInfo,
	}

	logHandler := slog.NewJSONHandler(os.Stdout, logOptions)
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	// Global random location generator
	// TODO: Read generator/cache/fetcher settings from config file
	locCircularLog := weather.NewLocCircularLog(1000, time.Minute*5)
	locGenerator := &weather.RandomLocationGenerator{
		ClientCount:    30,
		BufferSize:     1000,
		MaxFetchPerSec: 30,
		Fetcher: &weather.RandomLocationFetcher{
			MaxRecordAge:    time.Minute * 5,
			MinRetryDelay:   time.Millisecond * 100,
			MaxRetryDelay:   time.Second * 6,
			MaxRetries:      3,
			RandomLocAPIURL: "https://locations.patch3s.dev/api/random",
			PointsAPIURL:    "https://api.weather.gov/points/%v,%v",
			HTTPClient: &http.Client{
				Timeout: time.Second * 10,
			},
		},
		Cache: locCircularLog,
	}

	locDataChan := locGenerator.Generate(context.Background())

	// Configure Multiplexer/Server
	mux := http.NewServeMux()
	mux.Handle("GET /", middleware.LoggingMiddleware(handlers.GetWeatherHandler(locDataChan, time.Second*10, locCircularLog)))

	// TODO: Read address/port, timeouts from config file
	server := &http.Server{
		Addr:              ":8888",
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 2,
		ReadTimeout:       time.Second * 5,
		WriteTimeout:      time.Second * 20,
		IdleTimeout:       time.Second * 30,
	}

	logger.Info("Server started")
	err := server.ListenAndServe()

	if err != nil {
		logger.Error("Server crashed", slog.Any("Error", err))
	}

	// TODO: Gracefully shutdown
}

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
		Level: slog.LevelWarn,
	}

	logHandler := slog.NewJSONHandler(os.Stdout, logOptions)
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	/*
		ClientCount:     50,
		BufferSize:      1000,
		MaxFetchPerSec: 30,

		50%% in 9.6130 secs
		75%% in 9.8494 secs
		90%% in 10.0005 secs
		95%% in 10.0011 secs
		99%% in 10.0017 secs

		[200] 1793 responses
		[408] 207 responses
	*/
	// Global random location generator
	locGenerator := &weather.RandomLocationGenerator{
		// TODO: Read generator/fetcher settings from config file
		ClientCount:    50,
		BufferSize:     1000,
		MaxFetchPerSec: 30,
		Fetcher: &weather.RandomLocationFetcher{
			MinRetryDelay:   time.Millisecond * 100,
			MaxRetryDelay:   time.Second * 6,
			MaxRetries:      3,
			RandomLocAPIURL: "https://locations.patch3s.dev/api/random",
			PointsAPIURL:    "https://api.weather.gov/points/%v,%v",
			HTTPClient: &http.Client{
				Timeout: time.Second * 20,
			},
		},
	}

	locData := locGenerator.Generate(context.Background())

	// Configure Multiplexer/Server
	mux := http.NewServeMux()
	mux.Handle("GET /", middleware.LoggingMiddleware(handlers.GetWeatherHandler(locData, time.Second*10)))

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
}

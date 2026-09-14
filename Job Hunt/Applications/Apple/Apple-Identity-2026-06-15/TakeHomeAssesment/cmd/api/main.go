package main

import (
	"TakeHomeAssessment/internal/config"
	"TakeHomeAssessment/internal/handlers"
	"TakeHomeAssessment/internal/health"
	"TakeHomeAssessment/internal/resilience"
	"TakeHomeAssessment/internal/weather"
	"context"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {

	// This context will be canceled when the parent signals that the process is terminated
	ctx, shutdown := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer shutdown()

	config, err := config.LoadConfig()

	if err != nil {
		slog.Error("Error reading service configuration", slog.Any("Error", err))
		return
	}

	logger := SetupLogger(&config.Observability)

	RunProfilingTools(&config.Observability, logger)

	locDataChan, locCircularLog, locGenHealthData := StartWeatherGeneration(ctx, &config.WeatherGenerator)

	// Configure Multiplexer/Server
	wsConfig := config.WebServer
	mux := http.NewServeMux()
	mux.Handle("GET /", handlers.GetWeatherHandler(locDataChan, wsConfig.WriteTimeout, locCircularLog))
	mux.Handle("GET /stats", handlers.GetStatsHandler(locGenHealthData))

	server := &http.Server{
		Addr:              wsConfig.Address,
		Handler:           mux,
		ReadHeaderTimeout: wsConfig.ReadTimeout,
		ReadTimeout:       wsConfig.ReadTimeout,
		WriteTimeout:      wsConfig.WriteTimeout,
		IdleTimeout:       wsConfig.IddleTimeout,
	}

	logger.Info("Starting Web Server")
	err = server.ListenAndServe()

	if err != nil {
		logger.Error("Server crashed", slog.Any("Error", err))
	}

	// TODO: Gracefully shutdown
}

func SetupLogger(config *config.ObservabilityConfig) *slog.Logger {
	// Configure application logging
	logOptions := &slog.HandlerOptions{
		Level: config.LoggingLevel,
	}

	logHandler := slog.NewJSONHandler(os.Stdout, logOptions)
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
	return logger
}

func RunProfilingTools(config *config.ObservabilityConfig, logger *slog.Logger) {
	// PPROF server (Profiling tool)
	if config.EnableProfiling {
		go func() {
			logger.Info("Starting pprof server")

			// Configure profile sample rate of go routine blocking events
			runtime.SetBlockProfileRate(1)
			// Configure franction of mutex contention events that are reported
			runtime.SetMutexProfileFraction(10)
			if err := http.ListenAndServe(config.ProfilingAddress, nil); err != nil {
				logger.Error("Pprof server failed", slog.Any("Error", err))
			}
		}()
	}
}

func StartWeatherGeneration(ctx context.Context, config *config.WeatherGeneratorConfig) (<-chan weather.LocationInfo, weather.CircularLog, *health.LocationGenerationHealthData) {

	// Global random location generator
	locCircularLog := weather.NewLocCircularLog(config.BufferSize, config.MaxRecordAge)
	instrumentedRoundTripper := health.NewInstrumentedRoundTripper(&http.Transport{
		ForceAttemptHTTP2: true,
		// Avoid creating new TCP connections continously
		MaxIdleConns:        config.ClientCount,
		MaxIdleConnsPerHost: config.ClientCount,
		IdleConnTimeout:     config.RequestTimeout,
	})

	client := &http.Client{
		Timeout:   config.RequestTimeout,
		Transport: instrumentedRoundTripper,
	}

	circuitBreaker := resilience.NewLeakyCircuitBreaker(ctx, 10, config.MaxRetryDelay)
	fetcher := weather.NewRandomLocationFetcher(config, client, circuitBreaker)

	locGenerator := health.NewInstrumentedGenerator(weather.NewRandomLocationGenerator(config, fetcher, locCircularLog))
	locDataChan := locGenerator.Generate(ctx)

	return locDataChan, locCircularLog, &health.LocationGenerationHealthData{
		Generator:               locGenerator,
		OutgoingReqRoundTripper: instrumentedRoundTripper,
	}
}

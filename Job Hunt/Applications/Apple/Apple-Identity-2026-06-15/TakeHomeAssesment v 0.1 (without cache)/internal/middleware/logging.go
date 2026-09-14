package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		requestID := fmt.Sprintf("r_%v", startTime.UnixNano())
		requestLogger := slog.With(
			slog.String("requestID", requestID),
			slog.String("path", r.URL.Path),
		)

		ctxWithLogger := context.WithValue(r.Context(), "logger", requestLogger)
		r = r.WithContext(ctxWithLogger)

		nextHandler.ServeHTTP(w, r)

		requestLogger.Info("Request completed", slog.Duration("Duration", time.Since(startTime)))
	})
}

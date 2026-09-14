package weather

import (
	"TakeHomeAssessment/internal/config"
	"context"
	"errors"
	"log/slog"
	"sync"
	"time"
)

// This struct is responsible for producing the data [LocationInfo] we will serve in our endpoint in the background
// it will replenish the returned channel as soon as the data is read.
type RandomLocationGenerator struct {
	ClientCount    int
	BufferSize     int
	MaxFetchPerSec int
	Fetcher        LocationFetcher
	Cache          CircularLog
}

func NewRandomLocationGenerator(config *config.WeatherGeneratorConfig, fetcher LocationFetcher, cache CircularLog) *RandomLocationGenerator {
	return &RandomLocationGenerator{
		ClientCount:    config.ClientCount,
		BufferSize:     config.BufferSize,
		MaxFetchPerSec: config.MaxFetchPerSec,
		Fetcher:        fetcher,
		Cache:          cache,
	}
}

func (r *RandomLocationGenerator) Generate(ctx context.Context) <-chan LocationInfo {
	locations := make(chan LocationInfo, r.BufferSize)
	wg := sync.WaitGroup{}

	fetchPerSec := r.MaxFetchPerSec

	if fetchPerSec == 0 {
		// Avoid divide by zero
		fetchPerSec = 5
	}

	rateLimiter := time.NewTicker(time.Second / time.Duration(fetchPerSec))

	slog.Info("Started generating location info data.", slog.Int("ClientCount", r.ClientCount))

	wg.Add(r.ClientCount)
	for i := 0; i < r.ClientCount; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-rateLimiter.C:
					// Block until we can proceed
				case <-ctx.Done():
					slog.Info("Cancel signal received. Stopping client fetching location info.")
					return
				}

				res, err := r.Fetcher.GetLocation(ctx)
				if err != nil {
					if errors.Is(err, context.Canceled) {
						slog.Info("Cancel signal received. Stopping client fetching location info.")
						return
					}

					slog.Error("Error fething location information", slog.Any("Error", err))
					continue
				}

				// TODO: Make sure we sanitize/truncate the response from third party before including in our logs as is.
				slog.Debug("New location data fetched.", slog.String("Location", res.Name))

				// Save it to the circular log as a cache of the last elements generated
				r.Cache.Write(res)

				select {
				case locations <- res:
					// Successful
				case <-ctx.Done():
					slog.Info("Cancel signal received. Stopping client fetching location info.")
					return
				}
			}
		}()
	}

	// Cleanup
	go func() {
		wg.Wait()
		rateLimiter.Stop()
		close(locations)
		slog.Info("All clients fetching location data closed")
	}()

	return locations
}

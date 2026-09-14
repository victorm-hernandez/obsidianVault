package weather

import (
	"context"
	"math/rand/v2"
	"testing"
	"time"
)

type MockLocationFetcher struct {
	Result LocationInfo
	Error  error
}

func (m *MockLocationFetcher) GetLocation(ctx context.Context) (LocationInfo, error) {
	// Simulate delays
	time.Sleep(time.Duration(rand.IntN(100)+50) * time.Millisecond)

	return m.Result, m.Error
}

func TestRandomLocationGenerator_Generate(t *testing.T) {
	testValue := LocationInfo{
		Name:      "Barcelona",
		Longitude: 1.2,
		Latitude:  3.4,
		Created:   time.Now(),
	}

	mockFetcher := &MockLocationFetcher{
		Result: testValue,
	}

	testContext, cancelFunc := context.WithCancel(t.Context())
	locGenerator := &RandomLocationGenerator{
		ClientCount:    100,
		BufferSize:     200,
		Fetcher:        mockFetcher,
		MaxFetchPerSec: 500,
	}

	locations := locGenerator.Generate(testContext)

	for _ = range 100 {
		loc := <-locations

		if loc.Name != testValue.Name {
			t.Errorf("Unexpected name. Expected:%v, Current:%v", testValue.Name, loc.Name)
		}
	}

	// Verify that canceling the context closes the underlying channel.
	cancelFunc()

	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		select {
		case _, ok := <-locations:
			if !ok {
				// Closed as expected, we are done
				return
			}
		case <-time.After(100 * time.Millisecond):
		}
	}

	t.Errorf("Timed out waiting for channel to close")
}

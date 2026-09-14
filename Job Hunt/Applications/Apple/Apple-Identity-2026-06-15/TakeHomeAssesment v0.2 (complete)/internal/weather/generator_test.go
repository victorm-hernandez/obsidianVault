package weather

import (
	"context"
	"math/rand/v2"
	"sync"
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

type MockCircularLog struct {
	data  LocationInfo
	mutex sync.Mutex
}

func (m *MockCircularLog) Write(data LocationInfo) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data = data
}

func (m *MockCircularLog) Read() (LocationInfo, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.data, true
}

func TestRandomLocationGenerator_Generate(t *testing.T) {
	testValue := LocationInfo{
		Name:      "Barcelona",
		Longitude: 1.2,
		Latitude:  3.4,
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}

	mockFetcher := &MockLocationFetcher{
		Result: testValue,
	}

	mockLog := &MockCircularLog{}

	testContext, cancelFunc := context.WithCancel(t.Context())
	locGenerator := &RandomLocationGenerator{
		ClientCount:    100,
		BufferSize:     200,
		Fetcher:        mockFetcher,
		MaxFetchPerSec: 500,
		Cache:          mockLog,
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

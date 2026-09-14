package weather

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestRandomLocationFetcher_GetLocation(t *testing.T) {
	randomAPIPath := "/random"
	forecastAPIPath := "/forecast"
	pointAPIPathPrefix := "/point"
	pointAPIPathTemplate := "/point/%v,%v"

	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if r.URL.Path == randomAPIPath {
			w.Write([]byte(`{"locations":[{"name":"Norfolk","latitude":36.84681,"longitude":-76.28522}]}`))
			return
		}

		if strings.HasPrefix(r.URL.Path, pointAPIPathPrefix) {

			responseStr := fmt.Sprintf(`
            { 
                "fieldId1": "Some other value",
                "properties": 
                {
                    "forecast": "http://%v/forecast",
                    "otherField": "SomeValue"
                }
            }`, r.Host) // Make sure we return the current test server URL in the response

			w.Write([]byte(responseStr))
			return
		}

		if r.URL.Path == forecastAPIPath {
			w.Write([]byte(`
            {
                "type": "Feature",
                "properties": 
                {
                    "units": "us",
                    "periods": [
                        {
                            "shortForecast": "Haze",
                            "detailedForecast": "Haze. Mostly sunny, with a high near 96. Southwest wind around 8 mph."
                        },
                        {
                            "shortForecast": "Haze",
                            "detailedForecast": "Haze. Partly cloudy, with a low around 74. Southwest wind around 8 mph."
                        }
                    ]
                }
            }`))
			return
		}
	}))

	defer server.Close()

	fetcher := RandomLocationFetcher{
		MinRetryDelay:   time.Millisecond * 100,
		MaxRetryDelay:   time.Second * 6,
		MaxRetries:      3,
		MaxRecordAge:    time.Minute * 5,
		RandomLocAPIURL: fmt.Sprintf("%v%v", server.URL, randomAPIPath),
		PointsAPIURL:    fmt.Sprintf("%v%v", server.URL, pointAPIPathTemplate),
		HTTPClient: &http.Client{
			Timeout: time.Millisecond * 500,
		},
	}

	res, err := fetcher.GetLocation(t.Context())

	if err != nil {
		t.Errorf("Unexpected error while calling GetLocation. Error: %v", err.Error())
	}

	expectedForecastStr := "Haze. Mostly sunny, with a high near 96. Southwest wind around 8 mph."
	if res.Forecast != expectedForecastStr {
		t.Errorf("Unexpected forecast string. Received: %v, Expected: %v'", res.Forecast, expectedForecastStr)
	}
}

func TestRandomLocationFetcher_GetLocation_RetryLogic(t *testing.T) {
	randomAPIPath := "/random"
	pointAPIPathTemplate := "/point/%v,%v"
	requestCount := 0
	maxRetries := 3

	testCtxt, cancelFunc := context.WithCancel(t.Context())

	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++

		if r.URL.Path != randomAPIPath && requestCount < 4 {
			t.Error("The random API request was not retried as expected.")
		}

		if r.URL.Path == randomAPIPath {
			switch requestCount {
			// Retriable errors
			case 1:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`ERROR 500`))
			case 2:
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte(`ERROR 429`))
			case 3:
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"locations":[{"name":"Norfolk","latitude":36.84681,"longitude":-76.28522}]}`))
			}
			return
		}

		// Slow request to give time for cancelation to happen
		cancelFunc()
		time.Sleep(time.Millisecond * 400)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"locations":[{"name":"Norfolk","latitude":36.84681,"longitude":-76.28522}]}`))
	}))

	defer server.Close()

	fetcher := RandomLocationFetcher{
		MinRetryDelay:   time.Millisecond * 100,
		MaxRetryDelay:   time.Second * 6,
		MaxRetries:      maxRetries,
		MaxRecordAge:    time.Minute * 5,
		RandomLocAPIURL: fmt.Sprintf("%v%v", server.URL, randomAPIPath),
		PointsAPIURL:    fmt.Sprintf("%v%v", server.URL, pointAPIPathTemplate),
		HTTPClient: &http.Client{
			Timeout: time.Millisecond * 500,
		},
	}

	_, err := fetcher.GetLocation(testCtxt)

	if requestCount != 4 {
		t.Errorf("Unexpected number of requests. With the retries it should be 4. It is %v", requestCount)
	}

	if err == nil || !strings.Contains(err.Error(), "context canceled") {
		t.Errorf("No cancelation error returned after context canceled the request mid flight.")
	}
}

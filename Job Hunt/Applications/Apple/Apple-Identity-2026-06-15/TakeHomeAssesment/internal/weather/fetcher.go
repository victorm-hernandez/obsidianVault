package weather

import (
	"TakeHomeAssessment/internal/config"
	"TakeHomeAssessment/internal/resilience"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"time"
)

func NewRandomLocationFetcher(config *config.WeatherGeneratorConfig, httpClient *http.Client, circuitBreaker resilience.CircuitBreaker) *RandomLocationFetcher {
	return &RandomLocationFetcher{
		MaxRecordAge:    config.MaxRecordAge,
		MinRetryDelay:   config.MinRetryDelay,
		MaxRetryDelay:   config.MaxRetryDelay,
		MaxRetries:      config.MaxRetryCount,
		RandomLocAPIURL: config.RandomLocAPIURL,
		PointsAPIURL:    config.PointsAPIURL,
		HTTPClient:      httpClient,
		circuitBreaker:  circuitBreaker,
	}
}

// This struct is responsible for calling the 3 weather APIs in sequence to
// retrieve the forecast for a random location [LocationInfo]
type RandomLocationFetcher struct {
	MinRetryDelay   time.Duration
	MaxRetryDelay   time.Duration
	MaxRetries      int
	RandomLocAPIURL string
	PointsAPIURL    string
	HTTPClient      *http.Client
	MaxRecordAge    time.Duration
	circuitBreaker  resilience.CircuitBreaker
}

const userAgentString = "(RandomLocationWeather/1.0, victormh@gmail.com)"

func (r *RandomLocationFetcher) GetLocation(ctx context.Context) (LocationInfo, error) {
	rndLocRespBody, err := r.requestWithRetries(ctx, r.RandomLocAPIURL, "GET")

	if err != nil {

		return LocationInfo{}, fmt.Errorf("Error while calling the random location API, Error:%v", err)
	}

	defer rndLocRespBody.Close()

	var rndLocResponse struct {
		Locations []struct {
			Name      string
			Latitude  float32
			Longitude float32
		}
	}

	err = json.NewDecoder(rndLocRespBody).Decode(&rndLocResponse)

	if err != nil {
		return LocationInfo{}, fmt.Errorf("Error while deserializing the random location API response. Error:%v", err)
	}

	if len(rndLocResponse.Locations) != 1 ||
		rndLocResponse.Locations[0].Name == "" {
		return LocationInfo{}, fmt.Errorf("Unexpected response from random location API: %+v", rndLocResponse)
	}

	pointsURL := fmt.Sprintf(r.PointsAPIURL, rndLocResponse.Locations[0].Latitude, rndLocResponse.Locations[0].Longitude)

	pointsRespBody, err := r.requestWithRetries(ctx, pointsURL, "GET")

	if err != nil {
		return LocationInfo{}, fmt.Errorf("Error while calling the points API. Error:%v", err)
	}

	defer pointsRespBody.Close()

	var pointResponse struct {
		Properties struct {
			Forecast string `json:"forecast"`
		} `json:"properties"`
	}

	err = json.NewDecoder(pointsRespBody).Decode(&pointResponse)

	if err != nil {
		return LocationInfo{}, fmt.Errorf("Error while deserializing the point API response. Error:%v", err)
	}

	// TODO: Further validations of the response from the points API
	// - Make sure is a valid URL
	// - Make sure is in a allow list domain
	if pointResponse.Properties.Forecast == "" {
		return LocationInfo{}, fmt.Errorf("Unexpected response from points API: %+v", pointResponse)
	}

	forecastRespBody, err := r.requestWithRetries(ctx, pointResponse.Properties.Forecast, "GET")

	if err != nil {
		return LocationInfo{}, fmt.Errorf("Error while calling the forecast API. Error:%v", err)
	}

	defer forecastRespBody.Close()

	var forecastResponse struct {
		Properties struct {
			Periods []struct {
				DetailedForecast string `json:"detailedForecast"`
			} `json:"periods"`
		} `json:"properties"`
	}

	err = json.NewDecoder(forecastRespBody).Decode(&forecastResponse)

	if err != nil {
		return LocationInfo{}, fmt.Errorf("Error while deserializing the forecast API response. Error:%v", err)
	}

	// TODO: Additional checks for the response
	// - Sanitize the `DetailedForecast` to make sure only harmless strings are included (no HTML, JavaScript, Etc.)
	if len(forecastResponse.Properties.Periods) < 1 || forecastResponse.Properties.Periods[0].DetailedForecast == "" {
		return LocationInfo{}, fmt.Errorf("Unexpected response from forecast API: %+v", forecastResponse)
	}

	result := LocationInfo{
		Name:      rndLocResponse.Locations[0].Name,
		Latitude:  rndLocResponse.Locations[0].Latitude,
		Longitude: rndLocResponse.Locations[0].Longitude,
		Forecast:  forecastResponse.Properties.Periods[0].DetailedForecast,
		ExpiresAt: time.Now().Add(r.MaxRecordAge),
	}

	return result, nil
}

func (r *RandomLocationFetcher) requestWithRetries(ctx context.Context, url string, verb string) (io.ReadCloser, error) {
	var lastError error

	for i := 0; i < r.MaxRetries && r.circuitBreaker.IsAllowed(); i++ {
		request, err := http.NewRequestWithContext(ctx, verb, url, nil)

		if err != nil {
			return nil, err
		}

		request.Header.Set("User-Agent", userAgentString)

		response, err := r.HTTPClient.Do(request)

		// TODO: Add missing retriable Errors/Status codes
		if err != nil {
			lastError = err
			r.circuitBreaker.RecordError()

			// Timeout
			if os.IsTimeout(err) || errors.Is(err, context.DeadlineExceeded) {
				select {
				case <-ctx.Done():
					// Fail fast in case operation is canceled
					return nil, ctx.Err()
				case <-time.After(r.calcRetryDelay(i)):
					// Retry after delay
				}

				continue // Retriable errors
			}

			return nil, lastError // Non retriable errors
		}

		if response.StatusCode != http.StatusOK {
			retryAfter := convertRetryAfterHeader(response.Header.Get("Retry-After"))

			if retryAfter.IsZero() {
				r.circuitBreaker.RecordError()
			} else {
				r.circuitBreaker.RecordErrorWithRetryAfter(retryAfter)
			}

			lastError = fmt.Errorf("Unexpected HTTP Status Code %v while calling %v %v", response.StatusCode, verb, url)

			// Empty and close the body stream so the TCP connection can be re-used
			io.Copy(io.Discard, response.Body)
			response.Body.Close()

			// Retriable HTTP Codes
			if response.StatusCode == http.StatusInternalServerError ||
				response.StatusCode == http.StatusTooManyRequests ||
				response.StatusCode == http.StatusServiceUnavailable {

				select {
				case <-ctx.Done():
					// Fail fast in case operation is canceled
					return nil, ctx.Err()
				case <-time.After(r.calcRetryDelay(i)):
					// Retry after delay
				}
				continue
			}

			return nil, lastError // Non retriable HTTP code
		}

		return response.Body, nil
	}

	return nil, lastError
}

// Calculate exponential backoff with jitters
func (r *RandomLocationFetcher) calcRetryDelay(retryCount int) time.Duration {
	jitters := time.Duration(rand.IntN(150)) * time.Millisecond
	delay := r.MinRetryDelay + (time.Duration(math.Exp2(float64(retryCount))) * r.MinRetryDelay) + jitters

	if delay > r.MaxRetryDelay {
		return r.MaxRetryDelay
	}

	return delay
}

func convertRetryAfterHeader(headerValue string) time.Time {
	if headerValue == "" {
		return time.Time{} // Zero
	}

	// The header may contain an ammont of seconds to wait
	delayInSecs, err := strconv.Atoi(headerValue)

	if err == nil {
		return time.Now().Add(time.Second * time.Duration(delayInSecs))
	}

	// The header may a date after which we can try again
	var retryTime time.Time
	retryTime, err = http.ParseTime(headerValue)

	if err == nil {
		return retryTime
	}

	return time.Time{}
}

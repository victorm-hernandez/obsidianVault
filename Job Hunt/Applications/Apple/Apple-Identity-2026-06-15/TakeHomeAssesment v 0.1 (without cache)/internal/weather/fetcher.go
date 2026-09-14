package weather

import (
	"TakeHomeAssesment/internal/common"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

type RandomLocationFetcher struct {
	MinRetryDelay   time.Duration
	MaxRetryDelay   time.Duration
	MaxRetries      int
	RandomLocAPIURL string
	PointsAPIURL    string
	HTTPClient      *http.Client
}

const userAgentString = "(TakeHomeAssesment/1.0, victormh@gmail.com)"

func (r *RandomLocationFetcher) GetLocation(ctx context.Context) (LocationInfo, error) {
	rndLocRespBody, err := r.requestWithRetries(ctx, r.RandomLocAPIURL, "GET", r.HTTPClient)

	if err != nil {
		return LocationInfo{}, &common.CompoundError{
			Message: "Error while calling the random location API",
			Err:     err,
		}
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
		return LocationInfo{}, &common.CompoundError{
			Message: "Error while deserializing the random location API response.",
			Err:     err,
		}
	}

	if len(rndLocResponse.Locations) != 1 ||
		rndLocResponse.Locations[0].Name == "" {
		return LocationInfo{}, fmt.Errorf("Unexpected response from random location API: %+v", rndLocResponse)
	}

	pointsURL := fmt.Sprintf(r.PointsAPIURL, rndLocResponse.Locations[0].Latitude, rndLocResponse.Locations[0].Longitude)

	pointsRespBody, err := r.requestWithRetries(ctx, pointsURL, "GET", r.HTTPClient)

	if err != nil {
		return LocationInfo{}, &common.CompoundError{
			Message: "Error while calling the points API.",
			Err:     err,
		}
	}

	defer pointsRespBody.Close()

	var pointResponse struct {
		Properties struct {
			Forecast string `json:"forecast"`
		} `json:"properties"`
	}

	err = json.NewDecoder(pointsRespBody).Decode(&pointResponse)

	if err != nil {
		return LocationInfo{}, &common.CompoundError{
			Message: "Error while deserializing the point API response.",
			Err:     err,
		}
	}

	// TODO: Further validations of the response from the points API
	// - Make sure is a valid URL
	// - Make sure is in a whitelisted domain
	if pointResponse.Properties.Forecast == "" {
		return LocationInfo{}, fmt.Errorf("Unexpected response from points API: %+v", pointResponse)
	}

	forecastRespBody, err := r.requestWithRetries(ctx, pointResponse.Properties.Forecast, "GET", r.HTTPClient)

	if err != nil {
		return LocationInfo{}, &common.CompoundError{
			Message: "Error while calling the forecast API.",
			Err:     err,
		}
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
		return LocationInfo{}, &common.CompoundError{
			Message: "Error while deserializing the forecast API response.",
			Err:     err,
		}
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
		Created:   time.Now(),
	}

	return result, nil
}

func (r *RandomLocationFetcher) requestWithRetries(ctx context.Context, url string, verb string, client *http.Client) (io.ReadCloser, error) {
	var lastError error

	for i := 0; i < r.MaxRetries; i++ {
		request, err := http.NewRequestWithContext(ctx, verb, url, nil)

		if err != nil {
			return nil, err
		}

		request.Header.Set("User-Agent", userAgentString)

		response, err := client.Do(request)

		// TODO: Add missing retriable Errors/Status codes
		if err != nil {
			lastError = err

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
			lastError = fmt.Errorf("Unexpected HTTP Status Code %v while calling %v %v", response.StatusCode, verb, url)

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

package health

import (
	"TakeHomeAssessment/api"
	"net/http"
	"sync"
	"time"
)

type InstrumentedRoundTripper struct {
	next         http.RoundTripper
	mutex        sync.RWMutex
	RequestCount int64

	// Domain, { HttpStatusCode, Stats}
	statsPerDomain map[string]api.DomainStats
}

func NewInstrumentedRoundTripper(next http.RoundTripper) *InstrumentedRoundTripper {

	if next == nil {
		next = http.DefaultTransport
	}

	return &InstrumentedRoundTripper{
		next:           next,
		statsPerDomain: make(map[string]api.DomainStats, 10),
	}
}

func (i *InstrumentedRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	started := time.Now()
	resp, err := i.next.RoundTrip(req)

	if err != nil {
		return resp, err
	}

	duration := time.Since(started)
	domain := req.URL.Hostname()
	statusCode := resp.StatusCode

	i.mutex.Lock()
	defer i.mutex.Unlock()

	_, ok := i.statsPerDomain[domain]

	if !ok {
		i.statsPerDomain[domain] = api.DomainStats{
			DomainName:   domain,
			StatsPerCode: make(map[int]api.RequestStats),
		}
	}

	_, ok = i.statsPerDomain[domain].StatsPerCode[statusCode]

	if !ok {
		i.statsPerDomain[domain].StatsPerCode[statusCode] = api.RequestStats{
			RequestCount:    1,
			DurationCeiling: duration,
			DurationFloor:   duration,
			DurationAverage: duration,
		}
	} else {

		stats := i.statsPerDomain[domain].StatsPerCode[statusCode]
		stats.RequestCount++

		if duration < stats.DurationFloor {
			stats.DurationFloor = duration
		}

		if duration > stats.DurationCeiling {
			stats.DurationCeiling = duration
		}

		// Calculate the average
		diff := duration - stats.DurationAverage
		stats.DurationAverage = stats.DurationAverage + (diff / time.Duration(stats.RequestCount))

		i.statsPerDomain[domain].StatsPerCode[statusCode] = stats
	}

	return resp, err
}

func (i *InstrumentedRoundTripper) GetStats() map[string]api.DomainStats {

	i.mutex.RLock()
	defer i.mutex.RUnlock()

	newInstance := make(map[string]api.DomainStats, len(i.statsPerDomain))

	for domain, domainStats := range i.statsPerDomain {
		newInstance[domain] = api.DomainStats{
			StatsPerCode: make(map[int]api.RequestStats, len(domainStats.StatsPerCode)),
		}

		for httpCode, codeStats := range domainStats.StatsPerCode {
			newInstance[domain].StatsPerCode[httpCode] = codeStats
		}
	}

	return newInstance
}

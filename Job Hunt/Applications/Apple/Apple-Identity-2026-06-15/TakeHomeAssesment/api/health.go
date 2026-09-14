package api

import "time"

type HealthResponse struct {
	GeneratorStatus   GeneratorStatus        `json:"generatorStatus"`
	OutgoingReqStatus map[string]DomainStats `json:"outgoingReqStatistics"`
}

type GeneratorStatus struct {
	BootStartTime time.Time     `json:"bootStartTime"`
	BootDuration  time.Duration `json:"bootDurationMs"`
	ItemCount     int           `json:"itemCount"`
	Ready         bool          `json:"ready"`
}

type DomainStats struct {
	DomainName   string               `json:"domainName"`
	StatsPerCode map[int]RequestStats `json:"statsPerHttpCode"`
}

type RequestStats struct {
	RequestCount    int64         `json:"requestCount"`
	DurationCeiling time.Duration `json:"durationCeiling"`
	DurationFloor   time.Duration `json:"durationFloor"`
	DurationAverage time.Duration `json:"durationAverage"`
}

package weather

import (
	"context"
	"time"
)

type LocationInfo struct {
	Name      string
	Latitude  float32
	Longitude float32
	Forecast  string
	ExpiresAt time.Time
}

type LocationFetcher interface {
	GetLocation(ctx context.Context) (LocationInfo, error)
}

type LocationGenerator interface {
	Generate(ctx context.Context) <-chan LocationInfo
}

type CircularLog interface {
	Write(data LocationInfo)
	Read() (LocationInfo, bool)
}

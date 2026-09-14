package health

import (
	"TakeHomeAssessment/api"
	"TakeHomeAssessment/internal/weather"
	"context"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"
)

type InstrumentedGenerator struct {
	generator weather.LocationGenerator
	startTime time.Time
	mutex     sync.RWMutex

	// Status
	bootDuration time.Duration
	itemCount    atomic.Int64
	ready        bool
}

func NewInstrumentedGenerator(next weather.LocationGenerator) *InstrumentedGenerator {

	return &InstrumentedGenerator{
		generator: next,
	}
}

// Decorator patter/pipeline to instrument the response times of the generator and its proper initialization
func (i *InstrumentedGenerator) Generate(ctx context.Context) <-chan weather.LocationInfo {
	incomingDataChan := i.generator.Generate(ctx)
	bufferSize := cap(incomingDataChan)

	// We make the forward channel with as much room as the original to avoid a bottleneck
	// IMPORTANT: This is effectively duplicating the size of the buffer although divided in two channels.
	outgoingDataChan := make(chan weather.LocationInfo, bufferSize)

	i.startTime = time.Now()
	i.ready = false

	// TODO: handle the case where this function is called twice, that should return an error. This is designed to be called once per instance.

	go func() {
		defer close(outgoingDataChan)

		for {
			// Generation event
			// generationStarted = time.now
			select {
			case <-ctx.Done():
				slog.Debug("Received signal to cancel on InstrumentedGenerator.")
				return

			case data, ok := <-incomingDataChan:

				if !ok {
					// Channel closed
					return
				}

				// Reading data event
				// generationDuration = time.since(generationStarted)
				// lastRead = time.now
				// Metric: reads per second

				i.itemCount.Add(1)

				if i.itemCount.Load() == int64(bufferSize) {
					i.mutex.Lock()
					// Finished initializing the buffer capture how long it took
					i.bootDuration = time.Since(i.startTime)
					i.ready = true
					i.mutex.Unlock()
				}

				// servedStarted = time.now
				outgoingDataChan <- data
				// Served data (consumer) event
				// serveDuration = time.since(serveStarted)
				// lastServed = time.now
				// Metric: served per second

				// How many items remain on the buffer, is it zero?
				// currentBufferSize := len(outgoingDataChan)
			}
		}
	}()

	return outgoingDataChan
}

func (i *InstrumentedGenerator) GetStatus() api.GeneratorStatus {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	var bootDuration time.Duration

	// If still booting, calculate the time so far otherwise return saved value
	if i.ready {
		bootDuration = i.bootDuration
	} else {
		bootDuration = time.Since(i.startTime)
	}

	return api.GeneratorStatus{
		Ready:         i.ready,
		BootDuration:  bootDuration,
		BootStartTime: i.startTime,
		ItemCount:     int(i.itemCount.Load()),
	}
}

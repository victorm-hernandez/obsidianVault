package weather

import (
	"sync"
	"time"
)

// Data structure that persists the last X records [LocationInfo]
// if new data is provided beyond X, the oldest data is replaced.
// If the data is stale, it will be removed from the log.
type LocationInfoLog struct {
	logs       []LocationInfo
	mutex      sync.Mutex
	writeIndex int
	readIndex  int
	length     int
	maxLogAge  time.Duration
}

func NewLocationInfoLog(capacity int, maxLogAge time.Duration) *LocationInfoLog {
	instance := &LocationInfoLog{
		logs:      make([]LocationInfo, capacity), // Fixed size, initialized
		maxLogAge: maxLogAge,
	}

	return instance
}

func (l *LocationInfoLog) Log(data LocationInfo) {
	// TODO: validate location info is not stale before saving
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.logs[l.writeIndex] = data
	l.writeIndex++
	l.length++

	if l.length > len(l.logs) {
		l.length = len(l.logs)
	}

	if l.writeIndex >= len(l.logs) {
		l.writeIndex = 0
	}

	// If write catch up with read, bump read
	if l.writeIndex == l.readIndex {
		l.readIndex++
	}
}

func (l *LocationInfoLog) Read() (LocationInfo, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	for {
		if l.readIndex >= len(l.logs) {
			l.readIndex = 0
		}

		// There is no more data to read
		if l.length <= 0 {
			return LocationInfo{}, false
		}

		// Ignore stale data
		if time.Since(l.logs[l.readIndex].Created) > (l.maxLogAge) {
			l.readIndex++
			l.length--
			continue
		}

		toRead := l.readIndex
		l.readIndex++
		return l.logs[toRead], true
	}
}

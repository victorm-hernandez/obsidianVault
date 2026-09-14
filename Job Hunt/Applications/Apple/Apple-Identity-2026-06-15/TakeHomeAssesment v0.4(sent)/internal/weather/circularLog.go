package weather

import (
	"sync"
	"time"
)

// Slice that persists a fixed set of records [LocationInfo]
// The read operations are circular, data is continously read in cycles as long as it is not stale.
// The write operations are also circular, once the structure is full, the oldest data is replaced.
type LocCircularLog struct {
	logs       []LocationInfo
	mutex      sync.Mutex
	writeIndex int
	readIndex  int
	length     int
}

func NewLocCircularLog(capacity int, MaxLogAge time.Duration) *LocCircularLog {
	instance := &LocCircularLog{
		logs: make([]LocationInfo, capacity), // Fixed size, initialized
	}

	return instance
}

func (l *LocCircularLog) Write(data LocationInfo) {
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

func (l *LocCircularLog) Read() (LocationInfo, bool) {
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
		if l.logs[l.readIndex].ExpiresAt.Before(time.Now()) {
			l.readIndex++
			l.length--
			continue
		}

		toRead := l.readIndex
		l.readIndex++
		return l.logs[toRead], true
	}
}

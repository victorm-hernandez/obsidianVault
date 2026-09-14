package resilience

import (
	"context"
	"sync/atomic"
	"time"
)

type CircuitBreaker interface {
	IsAllowed() bool
	RecordError()
	RecordErrorWithRetryAfter(retryAfter time.Time)
	RecordSuccess()
}

const (
	statusClosed   int32 = 0
	statusOpened   int32 = 1
	statusHalfOpen int32 = 2
)

const (
	errorLeakRate   = 10
	errorLeakPeriod = time.Second
)

type LeakyCircuitBreaker struct {
	maxErrors       int32
	errorCount      atomic.Int32
	defaultCooldown time.Duration
	lastTripped     atomic.Int64 // time in unix epoch
	lastHealthProbe atomic.Int64 // time in unix epoch
	retryAfter      atomic.Int64 // time in unix epoch
	status          atomic.Int32
}

func NewLeakyCircuitBreaker(ctx context.Context, maxErrors int32, defaultCooldown time.Duration) *LeakyCircuitBreaker {
	instance := &LeakyCircuitBreaker{
		maxErrors:       maxErrors,
		defaultCooldown: defaultCooldown,
	}

	go instance.leakErrorsLoop(ctx)

	return instance
}

func (l *LeakyCircuitBreaker) leakErrorsLoop(ctx context.Context) {
	ticker := time.NewTicker(errorLeakPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// Service is shutting down
			return
		case <-ticker.C:
			// Leak (decrement) the error count over time
			l.errorCount.Add(errorLeakRate * (-1))

			if l.errorCount.Load() <= 0 {
				l.errorCount.Store(0)
			}
		}
	}
}

func (l *LeakyCircuitBreaker) RecordError() {
	// Zero time
	l.RecordErrorWithRetryAfter(time.Time{})
}

func (l *LeakyCircuitBreaker) RecordErrorWithRetryAfter(retryAfter time.Time) {

	l.errorCount.Add(1)

	status := l.status.Load()
	if status == statusOpened {
		// State is opened already, no need to do anything else
		return
	}

	var retryAfterUnix int64 = 0

	if !retryAfter.IsZero() {
		// Comes from response HTTP Header
		retryAfterUnix = retryAfter.UnixNano()
	}

	// Open circuit if either condition is met:
	// 1. The target service asked us explicitely to stop calling until this time
	// 2. If the max number of errors has been reached
	// 3. If there is an ongoing probe trying to close the circuit and failed
	if !retryAfter.IsZero() || l.errorCount.Load() > l.maxErrors || status == statusHalfOpen {

		// Compare and swap guarantees that only a single go routine will access this if statement
		if l.status.CompareAndSwap(statusClosed, statusOpened) || l.status.CompareAndSwap(statusHalfOpen, statusOpened) {
			l.lastTripped.Store(time.Now().UnixNano())

			defaultRetryAfterUnix := time.Now().Add(l.defaultCooldown).UnixNano()

			// Check if the default cooldown is longer than any explicit retry-after
			// and if it is prefer it
			if defaultRetryAfterUnix > retryAfterUnix {
				retryAfterUnix = defaultRetryAfterUnix
			}

			l.retryAfter.Store(retryAfterUnix)
		}
	}
}

func (l *LeakyCircuitBreaker) RecordSuccess() {
	l.errorCount.Store(0)
	l.status.CompareAndSwap(statusHalfOpen, statusClosed)
}

func (l *LeakyCircuitBreaker) IsAllowed() bool {
	status := l.status.Load()

	if status == statusClosed {
		// Regular requests
		return true
	}

	if status == statusHalfOpen {
		// Health probe in progress

		// If the health probe took longer than the default cooldown, try again
		retryProbeAfter := time.Unix(l.lastHealthProbe.Load(), 0).Add(l.defaultCooldown)

		if time.Now().After(retryProbeAfter) {
			// Setting to closed, so the next attempt creates a new health probe
			l.status.CompareAndSwap(statusHalfOpen, statusClosed)
		}

		return false
	}

	// Status opened, stop requests until retryAfter time
	if l.retryAfter.Load() < time.Now().UnixNano() {

		// Compare and swap acts as a gateway so only a single routine goes through
		if l.status.CompareAndSwap(statusOpened, statusHalfOpen) {
			l.lastHealthProbe.Store(time.Now().UnixNano())
			return true
		}
	}

	// Not time to retry, or someone else started the probe
	return false
}

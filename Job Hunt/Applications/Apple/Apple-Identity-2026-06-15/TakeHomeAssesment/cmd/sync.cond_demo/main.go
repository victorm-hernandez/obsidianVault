package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	// Our shared state
	var stopRunning atomic.Bool
	stopRunning.Store(true)

	// CONSUMER
	go func() {

		if stopRunning.Load() {

			cond.L.Lock()
			// Wait until there is a signal to continue
			for stopRunning.Load() {
				fmt.Println("Consumer: Waiting signal to process...")
				cond.Wait()
				// Wait() automatically unlocks 'mu' while suspended,
				// and re-locks 'mu' before proceeding.

			}
			cond.L.Unlock()

			// Condition met! We safely have the lock.
			fmt.Printf("Consumer: Processed item")
		}
	}()

	// PRODUCER
	time.Sleep(1 * time.Second) // Simulate time taken to produce

	cond.L.Lock()
	stopRunning.Store(false)
	fmt.Println("Producer: Enabled processing.")

	// Wake up a waiting consumer to check the condition
	cond.Signal()
	cond.L.Unlock()

	// Wait a moment for the consumer to finish printing
	time.Sleep(1 * time.Second)
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Metrics struct {
	TotalRequests int64
	ErrorCount    int64
}

func (m *Metrics) incrementRequests() {
	// m.TotalRequests++ // dangerous
	atomic.AddInt64(&m.TotalRequests, 1)
}

func (m *Metrics) incrementErrors() {
	// m.ErrorCount++ // dangerous
	atomic.AddInt64(&m.ErrorCount, 1)
}

func startAtomic() {
	fmt.Println("ATOMIC STARTING...")

	myMetrics := Metrics{}
	var wg sync.WaitGroup

	for x := 0; x < 1000; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			myMetrics.incrementRequests()
			if x != 0 && x%2 == 0 {
				myMetrics.incrementErrors()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Requests: %d, errors: %d\n", myMetrics.TotalRequests, myMetrics.ErrorCount)
	fmt.Printf("ATOMIC FINISHED.\n\n")
}

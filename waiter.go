package wait

import (
	"fmt"
	"time"
)

// Waiter holds the channels
type Waiter struct {
	requests chan string
	limiter  <-chan time.Time
}

// NewWaiter returns a pre-configured Waiter
func NewWaiter() (waiter Waiter) {
	w := Waiter{
		requests: make(chan string),
		limiter:  time.Tick(time.Second),
	}
	w.SetLimit(time.Second)
	go w.consume()
	return w
}

// StoreEvent processes the payload
func (w *Waiter) StoreEvent(payload string) {
	w.requests <- payload
}

// SetLimit sets a new rate limit
func (w *Waiter) SetLimit(limit time.Duration) {
	w.limiter = time.Tick(limit)
}

func (w *Waiter) consume() {
	for request := range w.requests {
		<-w.limiter
		// Request to external service goes here
		fmt.Printf("request: %s, received: %s\n", request, time.Now())
	}
}

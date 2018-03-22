package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter channel ticks every 200ms
	limiter := time.Tick(200 * time.Millisecond)

	// block on a receive from limiter channel
	// this limits to handling a request every 200ms
	for req := range requests {
		<-limiter // blocks until limiter receives a tick
		fmt.Println("request", req, time.Now())
	}

	// buffering the limiter channel allows us to handle short bursts of requests while preserving the overall rate limit
	burstyLimiter := make(chan time.Time, 3)
	// fill the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// every 200ms, try to add a new value to burstyLimiter
	go func() {
		for t := range time.Tick(200 * time.Microsecond) {
			burstyLimiter <- t
		}
	}()

	// simulate 5 more requests.  The first 3 will benefit from burstCapability now
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

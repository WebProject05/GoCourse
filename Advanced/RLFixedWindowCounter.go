package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiterFW struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewLimiter(limit int, window time.Duration) *RateLimiterFW {
	return &RateLimiterFW{
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiterFW) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}

	return false
}

func windowCounter() {
	limiter := NewLimiter(5, time.Second)

	var wg sync.WaitGroup

	// simulate 20 concurrent requests
	for i := 1; i <= 20; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			if limiter.Allow() {
				fmt.Println("Request", id, "Allowed")
			} else {
				fmt.Println("Request", id, "Denied")
			}
		}(i)
		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
}

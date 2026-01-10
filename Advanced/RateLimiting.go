package main

import (
	"fmt"
	"time"
)

/*
	Prevents Overload on API, abuse prevention and cost managment
	Controlls the throughput of a server, handling the number
	of requests the server can accept.
*/

// ==== TOKEN BUCKET ALGO ====

type RateLimiter struct {
	tokens chan struct{}
	refillTime time.Duration
}


func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}

	for range rateLimit {
		rl.tokens <- struct{}{}
	}
	go rl.startRefill()
	return rl
}

/*
	Without startRefill

	Timeline:

	Initial tokens = 30
	30 requests arrive → allowed
	Bucket empty
	ALL future requests → denied forever.



	With startRefill

	Timeline (with refillTime = 1s):
	Every 1 second, one token is added
	If bucket is full → token is dropped (default)
	If bucket has space → capacity restored

	This makes the limiter:
	Time-aware
	Self-healing
	Sustainable under load
*/
func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	for {
		select {
		case <- ticker.C:
			select {
			case rl.tokens <- struct{}{}:   // If there is space in bucket send a struct as a token request
			default:	// If the bucket if full do noting
			}
		}
	}
}

func (rl *RateLimiter) allow() bool {
	select {
	case<- rl.tokens:
		return true
	default:
		return false
	}
}


func main() {
	rateLimiter := NewRateLimiter(100, time.Second)
	count := 0

	for range 200 {
		if rateLimiter.allow() {
			fmt.Println("Request Allowed")
			count++
		} else {
			fmt.Println("Request Denied")
		}
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("Final accepted count:", count)
}
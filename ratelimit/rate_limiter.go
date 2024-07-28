package ratelimit

import (
	"sync"
	"time"
)

// NOTE: Due to lack of specification, I took the liberty to assume that the rate-limiter is a global one,
// and should not be applied per unique source nor type of request.
type RateLimiter struct {
	rateRequests int
	rateInterval time.Duration
	lastTime     time.Time
	mutex        sync.Mutex
}

func NewRateLimiter(rateRequests int, rateInterval time.Duration) *RateLimiter {
	return &RateLimiter{
		rateRequests: rateRequests,
		rateInterval: rateInterval,
		lastTime:     time.Now().Add(-rateInterval), // Initialize to allow first request
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	if now.Sub(rl.lastTime) < rl.rateInterval {
		return false
	}

	rl.lastTime = now
	return true
}

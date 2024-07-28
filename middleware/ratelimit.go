package middleware

import (
	"Torq_IPGeo_Assignment/ratelimit"
	"log"
	"net/http"
)

// RateLimiterHandler is a middleware to enforce Rate Limit on middleware runtime
type RateLimiter struct {
	limiter *ratelimit.RateLimiter
}

func NewRateLimiter(rl *ratelimit.RateLimiter) RateLimiter {
	return RateLimiter{limiter: rl}
}

func (rl *RateLimiter) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.limiter.Allow() {
			log.Println("Rate limit exceeded")
			http.Error(w, `{"error": "rate limit exceeded"}`, http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

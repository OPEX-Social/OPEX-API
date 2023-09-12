// Version 1.0
package main

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// RateLimiter is a middleware that rate limits requests based on the client's IP address and route.
type RateLimiter struct {
	limiterMap map[string]map[string]*rate.Limiter
	mu         sync.Mutex
}

// NewRateLimiter creates a new RateLimiter instance.
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limiterMap: make(map[string]map[string]*rate.Limiter),
	}
}

// Middleware applies rate limiting to incoming requests based on the client's IP address and route.
func (rl *RateLimiter) RateLimit(route string, limit int, window time.Duration, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)

		rl.mu.Lock()
		routeLimiterMap, ok := rl.limiterMap[route]
		if !ok {
			routeLimiterMap = make(map[string]*rate.Limiter)
			rl.limiterMap[route] = routeLimiterMap
		}

		limiter, ok := routeLimiterMap[clientIP]
		if !ok {
			limiter = rate.NewLimiter(rate.Limit(limit), int(window.Seconds()))
			routeLimiterMap[clientIP] = limiter
		}
		rl.mu.Unlock()

		if limiter.Allow() {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		}
	})
}

package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	limiter := NewRateLimiter(10, 5*time.Second)

	http.HandleFunc("/", rateLimitMiddleware(limiter, indexHandler))
	if err := http.ListenAndServe("localhost:6789", nil); err != nil {
		log.Printf("server error: %s", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong\n"))
}

type rateLimiter struct {
	AllowedRequestsNum int
	Window             time.Duration

	mu        sync.Mutex
	counter   int
	lastReset time.Time
}

func NewRateLimiter(requests int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		AllowedRequestsNum: requests,
		Window:             window,
		lastReset:          time.Now(),
	}
}

func (rl *rateLimiter) Inc() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.Sub(rl.lastReset) > rl.Window {
		rl.counter = 0
		rl.lastReset = now
	}

	rl.counter += 1
}

func (rl *rateLimiter) IsLimitExceeded() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.Sub(rl.lastReset) > rl.Window {
		rl.counter = 0
		rl.lastReset = now
	}

	return rl.counter > rl.AllowedRequestsNum
}

func rateLimitMiddleware(limiter *rateLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request from", r.RemoteAddr)

		if limiter.IsLimitExceeded() {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		limiter.Inc()

		next(w, r)
	}
}

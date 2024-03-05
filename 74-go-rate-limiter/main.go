package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	limiter := &rateLimiter{
		AllowedRequestsNum: 10,
	}

	http.HandleFunc("GET /", rateLimitMiddleware(limiter, indexHandler))
	if err := http.ListenAndServe("localhost:6789", nil); err != nil {
		log.Printf("server error: %s", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong\n"))
}

type rateLimiter struct {
	AllowedRequestsNum int

	mu      sync.Mutex
	counter int
}

func (rl *rateLimiter) Inc() {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.counter += 1
}

func (rl *rateLimiter) Clear() {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.counter = 0
}

func (rl *rateLimiter) IsLimitExceeded() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	return rl.counter > rl.AllowedRequestsNum
}

func rateLimitMiddleware(limiter *rateLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("perform request")

		if limiter.IsLimitExceeded() {
			go func() {
				<-time.After(200 * time.Millisecond)
				log.Println("clear limiter")
				limiter.Clear()
			}()

			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		limiter.Inc()

		next(w, r)
	}
}

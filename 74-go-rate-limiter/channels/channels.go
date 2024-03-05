package main

import (
	"log"
	"net/http"
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
	tickets    chan time.Time
	timeWindow time.Duration
}

func NewRateLimiter(requests int, window time.Duration) *rateLimiter {
	limiter := &rateLimiter{
		tickets:    make(chan time.Time, requests),
		timeWindow: window,
	}

	go func() {
		tick := time.NewTicker(limiter.timeWindow / time.Duration(requests))
		for t := range tick.C {
			limiter.tickets <- t
		}
	}()

	return limiter
}

func (rl *rateLimiter) Allow() bool {
	select {
	case <-rl.tickets:
		return true
	default:
		return false
	}
}

func rateLimitMiddleware(limiter *rateLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("perform request")

		if !limiter.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		next(w, r)
	}
}

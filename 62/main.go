package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(100 * time.Millisecond))

	r.Get("/", handler)

	http.ListenAndServe(":3333", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()

	time.Sleep(500 * time.Millisecond)

	// select {
	// default:
	// case <-ctx.Done():
	// 	return
	// }

	w.Write([]byte("hello world"))
}

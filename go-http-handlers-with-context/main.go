package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

const ContextKey = "my-data"

type Config struct {
	Host, Port, Message string
}

func Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	msg := r.Context().Value(ContextKey).(string)
	w.Write([]byte(fmt.Sprintf("Hello, %s!\nMsg: %s\n", vars["name"], msg)))
}

func WithConfig(c Config, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ContextKey, c.Message)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	var cfg Config
	flag.StringVar(&cfg.Host, "host", "0.0.0.0", "HTTP host")
	flag.StringVar(&cfg.Port, "port", "8000", "HTTP port")
	flag.StringVar(&cfg.Message, "message", "secret!", "Custom message to pass")
	flag.Parse()

	r := mux.NewRouter()
	// Pass config to handler with wrapper
	r.HandleFunc("/{name}", WithConfig(cfg, http.HandlerFunc(Handle)))
	// Bind to a port and pass our router in
	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	log.Printf("listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

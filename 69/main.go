package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"log/slog"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	slog.Debug("debug message", "key", "value")

	ctx = context.WithValue(ctx, "key", "new value")
	slog.DebugContext(ctx, "debug message", "key", "value")
}

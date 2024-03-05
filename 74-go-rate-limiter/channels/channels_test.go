package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	limiter := NewRateLimiter(3, 1*time.Second)
	handler := rateLimitMiddleware(limiter, indexHandler)

	server := httptest.NewServer(handler)
	defer server.Close()

	client := server.Client()

	// Make 3 requests in quick succession
	for i := 0; i < 3; i++ {
		resp, err := client.Get(server.URL)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// The first 2 requests should succeed
		if i < 2 && resp.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}

		// The third request should be rate limited
		if i == 3 && resp.StatusCode != http.StatusTooManyRequests {
			t.Errorf("expected status 429, got %d", resp.StatusCode)
		}

		resp.Body.Close()
	}

	// Wait for the rate limit to reset
	time.Sleep(1 * time.Second)

	// The next request should succeed
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()
}

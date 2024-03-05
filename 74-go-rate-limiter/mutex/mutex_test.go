package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	limiter := NewRateLimiter(3, time.Second)
	handler := rateLimitMiddleware(limiter, indexHandler)

	// Create a new request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Make 3 requests in quick succession
	for i := 0; i < 3; i++ {
		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		// Call the handler function
		handler.ServeHTTP(rr, req)

		// For the first 2 requests, the status code should be 200
		if i < 2 && rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		}

		// For the 3rd request, the status code should be 429 (Too Many Requests)
		if i == 3 && rr.Code != http.StatusTooManyRequests {
			t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusTooManyRequests)
		}
	}

	// Wait for more than a second
	time.Sleep(time.Second + time.Millisecond*100)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler.ServeHTTP(rr, req)

	// The status code should be 200 again, because the rate limit should have been reset
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

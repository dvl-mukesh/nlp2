package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	// "github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	s := Server{logger: log.Default()}
	// Note: This bypasses routing & middleware
	s.healthHandler(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %d, actual %d", http.StatusOK, resp.StatusCode)
	}
}

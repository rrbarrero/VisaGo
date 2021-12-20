package main

import (
	"net/http"
	"net/http/httptest"
	"rbarrero/visago/routes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPing check service health
func TestPing(t *testing.T) {
	router := routes.SetupRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "pong")

}

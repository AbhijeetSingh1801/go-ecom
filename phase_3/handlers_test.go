package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AbhijeetSingh1801/go-ecom/phase_3/routes"
	"github.com/gin-gonic/gin"
)

func TestHealthCheck(t *testing.T) {
	router := gin.New()
	routes.SetupRoutes(router, nil)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

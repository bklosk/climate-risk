package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCoords200(t *testing.T) {
	// Create a new Gin router
	var router = gin.Default()
	router.GET("/coords/:address", getCoords)

	// Create a new HTTP recorder to capture the response
	var recorder = httptest.NewRecorder()
	// Create a new HTTP request to the /coords/:address endpoint
	req, err := http.NewRequest("GET", "/coords/2426%20ontario%20rd%20NW", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Serve the HTTP request to the recorder
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

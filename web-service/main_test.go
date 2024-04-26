package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetBen(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Register the getBen handler
	router.GET("/ben", getBen)

	// Create a new HTTP request to the /ben endpoint
	req, err := http.NewRequest("GET", "/ben", nil)
	if (err != nil) {
		t.Fatalf("An error occurred while creating a request: %v", err)
	}

	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()

	// Serve the HTTP request to the recorder
	router.ServeHTTP(recorder, req)

	// Assert that the response status code is 200 OK
	if (recorder.Code != http.StatusOK) {
		t.Fatalf("Expected status code %v but got %v", http.StatusOK, recorder.Code)
	}

	// Assert that the response body contains the expected JSON data
	expectedName := "Ben"
	var actual user
	json.Unmarshal(recorder.Body.Bytes(), &actual)

	if (expectedName != actual.Name) {
		t.Fatalf("Expected JSON %v but got %v", expectedName, recorder.Body.String())
	}
}

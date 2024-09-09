package consents

import (
	// for testing the functions
	"bytes"
	"go-assigment/data"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestCheckConsentWithConsentGiven(t *testing.T) {
	//gin.SetMode(gin.TestMode)
	data.Log = logrus.New()
	// Create a new Gin router
	router := gin.Default()

	// Define a test route for the CheckConsent handler
	router.POST("/consents/:dialogID", CheckConsent)

	// Create a test request
	dialogID := "123"
	message := "true"
	requestBody := []byte(message)

	req, err := http.NewRequest("POST", "/consents/"+dialogID, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create test request: %s", err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Use the router to handle the test request
	router.ServeHTTP(recorder, req)

	// Verify the response
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	responseBody := recorder.Body.String()
	expectedResponse := `"Dialog data saved"`
	if responseBody != expectedResponse {
		t.Errorf("Expected response body to be '%s', but got '%s'", expectedResponse, responseBody)
	}
}

func TestCheckConsentWithoutConsent(t *testing.T) {
	//gin.SetMode(gin.TestMode)
	data.Log = logrus.New()
	// Create a new Gin router
	router := gin.Default()

	// Define a test route for the CheckConsent handler
	router.POST("/consents/:dialogID", CheckConsent)

	// Create a test request
	dialogID := "123"
	message := "false"
	requestBody := []byte(message)

	req, err := http.NewRequest("POST", "/consents/"+dialogID, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create test request: %s", err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Use the router to handle the test request
	router.ServeHTTP(recorder, req)

	// Verify the response
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	responseBody := recorder.Body.String()
	expectedResponse := `"Dialog data deleted"`
	if responseBody != expectedResponse {
		t.Errorf("Expected response body to be '%s', but got '%s'", expectedResponse, responseBody)
	}
}

func TestCheckConsentWithInvalidPayload(t *testing.T) {
	//gin.SetMode(gin.TestMode)
	data.Log = logrus.New()
	// Create a new Gin router
	router := gin.Default()

	// Define a test route for the CheckConsent handler
	router.POST("/consents/:dialogID", CheckConsent)

	// Create a test request
	dialogID := "123"
	message := "random"
	requestBody := []byte(message)

	req, err := http.NewRequest("POST", "/consents/"+dialogID, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create test request: %s", err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Use the router to handle the test request
	router.ServeHTTP(recorder, req)

	// Verify the response
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
	}

	responseBody := recorder.Body.String()
	expectedResponse := `"Invalid request body"`
	if responseBody != expectedResponse {
		t.Errorf("Expected response body to be '%s', but got '%s'", expectedResponse, responseBody)
	}
}

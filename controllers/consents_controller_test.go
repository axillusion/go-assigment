package controllers

import (
	// for testing the functions
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/axillusion/go-assigment/commons"
	"github.com/axillusion/go-assigment/logger"
	"github.com/axillusion/go-assigment/mocks"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/mock"
)

var router *gin.Engine
var service mocks.MockDataService

func setupConsent() *gin.Engine {
	commons.Log = logger.NewLogger()
	service := new(mocks.MockDataService)
	service.On("ModifyDB", "123", mock.AnythingOfType("chan<- error")).Return(nil).Once()
	service.On("Delete", "123", mock.AnythingOfType("chan<- error")).Return(nil).Once()
	consentController := NewConsentController(service)
	router := gin.Default()
	router.POST("/consents/:dialogID", consentController.CheckConsent)
	return router
}

func TestCheckConsentWithConsentGiven(t *testing.T) {
	router = setupConsent()
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
	router = setupConsent()
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
	router = setupConsent()
	dialogID := "123"
	message := "random message that will fail"
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

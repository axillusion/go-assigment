package tests

import (
	// for testing the functions
	"bytes"
	"main/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestCheckConsentWithConsentGiven(t *testing.T) {
	//gin.SetMode(gin.TestMode)
	router, service := utils.InitServerTestWithMock()

	// Create a test request
	dialogID := "123"
	message := "true"
	requestBody := []byte(message)

	service.On("ModifyDB", dialogID, mock.AnythingOfType("chan<- error")).Return(nil).Once()
	service.On("Delete", dialogID, mock.AnythingOfType("chan<- error")).Return(nil).Once()
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
	router, service := utils.InitServerTestWithMock()

	// Create a test request
	dialogID := "123"
	message := "false"
	requestBody := []byte(message)

	service.On("ModifyDB", dialogID, mock.AnythingOfType("chan<- error")).Return(nil).Once()
	service.On("Delete", dialogID, mock.AnythingOfType("chan<- error")).Return(nil).Once()

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
	router, service := utils.InitServerTestWithMock()

	// Create a test request
	dialogID := "123"
	message := "random message that will fail"
	requestBody := []byte(message)

	service.On("ModifyDB", dialogID, mock.AnythingOfType("chan<- error")).Return(nil).Once()
	service.On("Delete", dialogID, mock.AnythingOfType("chan<- error")).Return(nil).Once()

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

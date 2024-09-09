package consents

import (
	// for testing the functions

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCheckConsent(t *testing.T) {
	// Create a new router
	router := gin.Default()

	// Register the CheckConsent endpoint
	router.POST("/check-consent/:dialogID", CheckConsent)

	// Create a new HTTP request with the desired method, URL, and request body
	reqBody := "true" // Set the request body as "true" for testing
	req, err := http.NewRequest("POST", "/check-consent/dialog123", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	res := httptest.NewRecorder()

	// Process the request using the router
	router.ServeHTTP(res, req)

	// Check the response status code
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %v, but got %v", http.StatusOK, res.Code)
	}

	// Check the response body
	expectedBody := "Dialog data saved"
	actualBody := res.Body.String()
	if actualBody != expectedBody {
		t.Errorf("Expected response body %q, but got %q", expectedBody, actualBody)
	}
}

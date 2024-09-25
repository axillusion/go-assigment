package tests

import (
	"bytes"
	"main/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestAddMessageToDialogInvalidBody(t *testing.T) {
	router, service := utils.InitServerTestWithMock()
	service.On("SaveToDB", "404", "404", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("chan<- error")).Return(nil).Once()
	// Create a test request
	requestBody := []byte(`this will fail`)
	req, err := http.NewRequest("POST", "/data/404/404", bytes.NewBuffer(requestBody))
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
}

func TestAddMessageToDialogValidBody(t *testing.T) {
	router, service := utils.InitServerTestWithMock()
	service.On("SaveToDB", "404", "404", "Test text", "EN", mock.AnythingOfType("chan<- error")).Return(nil).Once()
	// Create a test request
	requestBody := []byte(`{"text": "Test text", "language": "EN"}`)
	req, err := http.NewRequest("POST", "/data/404/404", bytes.NewBuffer(requestBody))
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
}

/*
func TestGetUserData(t *testing.T) {
	Init()
	r := gin.Default()
	r.GET("/data", GetUserData)
	SaveToDB(testDialogData.CustomerID, testDialogData.DialogID, testDialogData.Text, testDialogData.Language)
	type testCase struct {
		name       string
		query      string
		statusCode int
	}

	testCases := []testCase{
		{
			name:       "With language and customerID",
			query:      "?language=EN&customerID=404",
			statusCode: http.StatusOK,
		},
		{
			name:       "With page and pageSize",
			query:      "?page=1&pageSize=10",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/data"+tc.query, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			r.ServeHTTP(w, req)

			if w.Code != tc.statusCode {
				t.Errorf("Expected status code %d but got %d", tc.statusCode, w.Code)
			}

			//expectedPayload := []DialogRow{testDialogData}
			var responseData []DialogRow

			err = json.Unmarshal(w.Body.Bytes(), &responseData)
			if err != nil {
				t.Fatalf("Failed to unmarshal response body: %v", err)
			}

			// Assert that the response has the expected HTTP status code
			if w.Code != tc.statusCode {
				t.Errorf("Expected status code %d but got %d", tc.statusCode, w.Code)
			}

			// Assert that the response body matches the expected payload
			if !reflect.DeepEqual(responseData, expectedPayload) {
				t.Errorf("Response body does not match expected payload.\nExpected: %v\nGot: %v", expectedPayload, responseData)
			}
		})
	}
	DeleteDialogData(testDialogData.DialogID)
}*/

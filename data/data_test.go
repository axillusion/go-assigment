package data

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var testDialogData = DialogRow{CustomerID: "404", DialogID: "404", Text: "Test text", Language: "EN"}

func TestSaveToDB(t *testing.T) {
	Init()

	SaveToDB(testDialogData.CustomerID, testDialogData.DialogID, testDialogData.Text, testDialogData.Language)

	query := Db.Select("dialogID", "customerID", "stext", "language").Where("dialogID = ?", testDialogData.DialogID)
	var entry DialogRow
	query.First(&entry)
	if entry.DialogID == "" {
		t.Errorf("Failed to add to the database")
	}
	DeleteDialogData(testDialogData.DialogID)
}

func TestDeleteDialogData(t *testing.T) {
	Init()
	SaveToDB(testDialogData.CustomerID, testDialogData.DialogID, testDialogData.Text, testDialogData.Language)
	DeleteDialogData(testDialogData.DialogID)
	query := Db.Select("dialogID", "customerID", "stext", "language").Where("dialogID = ?", testDialogData.DialogID)
	var entry DialogRow
	query.First(&entry)
	if entry.DialogID != "" {
		t.Errorf("Failed to delete from the database")
	}
}

func TestAddMessageToDialogInvalidBody(t *testing.T) {
	Init()
	router := gin.Default()
	// Define a test route for the CheckConsent handler
	router.POST("/data/:customerID/:dialogID", AddMessageToDialog)
	requestBody := []byte("")
	req, err := http.NewRequest("POST", "/data/"+testDialogData.CustomerID+"/"+testDialogData.DialogID, bytes.NewBuffer(requestBody))
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
	expectedResponse := `"Invalid JSON payload"`
	if responseBody != expectedResponse {
		t.Errorf("Expected response body to be '%s', but got '%s'", expectedResponse, responseBody)
	}
}

func TestAddMessageToDialogValidBody(t *testing.T) {
	Init()
	router := gin.Default()
	// Define a test route for the CheckConsent handler
	router.POST("/data/:customerID/:dialogID", AddMessageToDialog)
	requestBody := Data{Text: testDialogData.Text, Language: testDialogData.Language}
	jsonPayload, err := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/data/"+testDialogData.CustomerID+"/"+testDialogData.DialogID, bytes.NewBuffer(jsonPayload))
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
	expectedResponse := `"Message saved"`
	if responseBody != expectedResponse {
		t.Errorf("Expected response body to be '%s', but got '%s'", expectedResponse, responseBody)
	}
}

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
			/*if !reflect.DeepEqual(responseData, expectedPayload) {
				t.Errorf("Response body does not match expected payload.\nExpected: %v\nGot: %v", expectedPayload, responseData)
			}*/
		})
	}
	DeleteDialogData(testDialogData.DialogID)
}

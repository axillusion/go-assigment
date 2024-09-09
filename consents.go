package main

import (
	// Import to handle HTTP messages
	"net/http"
	// Import for using api endpoints
	"github.com/gin-gonic/gin"
	// For text processing
	"io"
)

// Endpoint to save dialog data based on users consent
func CheckConsent(c *gin.Context) {
	// The ID of the dialog
	dialogID := c.Param("dialogID")
	// The request body should contain the message "true" or "false"
	body, error := io.ReadAll(c.Request.Body)
	if error != nil {
		c.JSON(http.StatusInternalServerError, "Failed to read request body")
		return
	}

	answer := string(body)
	// Returns a bad request in case the body does not contain just the message "true" or "false"
	if answer != "true" && answer != "false" {
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}

	if answer == "false" {
		c.JSON(http.StatusOK, "Dialog data deleted")
		// deletes the dialog data if the user does not consent
		DeleteDialogData(dialogID)
	} else {
		c.JSON(http.StatusOK, "Dialog data saved")
	}
}

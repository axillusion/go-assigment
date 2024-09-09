package consents

import (
	// Import to handle HTTP messages

	"net/http"
	// Import for using api endpoints
	"github.com/gin-gonic/gin"

	// Import for logging
	"github.com/sirupsen/logrus"

	// For text processing
	"io"

	"go-assigment/data"
)

// Endpoint to save dialog data based on users consent
func CheckConsent(c *gin.Context) {
	// The ID of the dialog
	dialogID := c.Param("dialogID")
	data.Log.WithFields(logrus.Fields{"dialogID": dialogID}).Info("Endpoint on /data called with on this dialogID")
	// The request body should contain the message "true" or "false"
	body, error := io.ReadAll(c.Request.Body)
	if error != nil {
		data.Log.Fatal("Failed to read request body")
		c.JSON(http.StatusInternalServerError, "Failed to read request body")
		return
	}

	answer := string(body)
	// Returns a bad request in case the body does not contain just the message "true" or "false"
	data.Log.WithFields(logrus.Fields{"body": answer}).Info("The body of the request was read")
	if answer != "true" && answer != "false" {
		data.Log.Warn("The body of the request is invalid, request wont be processed")
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}

	data.Log.Info("Processing consent request")
	if answer == "false" {
		c.JSON(http.StatusOK, "Dialog data deleted")
		// deletes the dialog data if the user does not consent
		defer data.DeleteDialogData(dialogID)
	} else {
		c.JSON(http.StatusOK, "Dialog data saved")
	}
}

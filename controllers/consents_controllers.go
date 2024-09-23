package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"

	"io"
	"main/commons"
	"main/services"
)

// Endpoint to save dialog data based on users consent
func CheckConsent(c *gin.Context) {
	// The ID of the dialog
	dialogID := c.Param("dialogID")
	commons.Log.WithFields(logrus.Fields{"dialogID": dialogID}).Info("Endpoint on /data called with on this dialogID")
	// The request body should contain the message "true" or "false"
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		commons.Log.Fatal("Failed to read request body")
		c.JSON(http.StatusInternalServerError, "Failed to read request body")
		return
	}

	answer := string(body)
	// Returns a bad request in case the body does not contain just the message "true" or "false"
	commons.Log.WithFields(logrus.Fields{"body": answer}).Info("The body of the request was read")
	if answer != "true" && answer != "false" {
		commons.Log.Warn("The body of the request is invalid, request wont be processed")
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}

	commons.Log.Info("Processing consent request")
	if answer == "false" {
		c.JSON(http.StatusOK, "Dialog data deleted")
		// deletes the dialog data if the user does not consent
		errorChannel := make(chan error)
		go services.DeleteDialogData(dialogID, errorChannel)
		err := <-errorChannel
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Could not delete dialog entries with this ID")
		}
	} else {
		c.JSON(http.StatusOK, "Dialog data saved")
	}
}

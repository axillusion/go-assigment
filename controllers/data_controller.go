package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"main/commons"
	"main/models"
	"main/services"
)

// Endpoint to add a new dialog message
func AddMessageToDialog(c *gin.Context) {
	customerID := c.Param("customerID")
	dialogID := c.Param("dialogID")
	commons.Log.WithFields(logrus.Fields{"customerID": customerID, "dialogID": dialogID}).Info("Post request for a new dialog entry with this params has been called")
	var messageContent models.Data
	err := c.BindJSON(&messageContent)
	// Message has to contain the text and the language of the message
	if err != nil {
		// Returns a bad request if the body is not valid
		c.JSON(http.StatusBadRequest, "Invalid JSON payload")
		commons.Log.Warn("The JSON payload of the request invalid, processing finished")
		return
	}
	commons.Log.Info("Processing request")

	errorChannel := make(chan error)
	go services.SaveToDB(customerID, dialogID, messageContent.Text, messageContent.Language, errorChannel)
	err = <-errorChannel
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to add to the database")
		return
	}
	c.JSON(http.StatusOK, "Message saved")
}

// Endpoint to fetch all the data
func GetUserData(c *gin.Context) {
	language := c.Query("language")
	customerID := c.Query("customerID")
	var dataPoints []map[string]interface{}

	query := commons.Db.Table("dialog_rows").
		Select("dialogID", "customerID", "stext", "language").
		Order("created_at DESC").
		Where("consent = ?", true)

	// Check if language argument is present
	if language != "" {
		query = query.Where("language = ?", language)
	}

	// Check if customerID argument is present
	if customerID != "" {
		query = query.Where("customerID = ?", customerID)
	}

	// Get the page number and page size from the query parameters
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	// Calculate the offset based on the page number and page size
	offset := (page - 1) * pageSize

	commons.Log.WithFields(logrus.Fields{"language": query.RowsAffected, "customerID": customerID, "page": page, "pageSize": pageSize}).Info("Received request to fetch data on data/ with the following parameters")
	if pageSize > 0 {
		query = query.Limit(pageSize).Offset(offset)
	}
	// Execute the query and scan the results into the dataPoints slice
	errorChannel := make(chan error)
	go services.FetchData(query, &dataPoints, errorChannel)
	err := <-errorChannel
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	c.IndentedJSON(http.StatusOK, dataPoints)
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/axillusion/go-assigment/commons"
	"github.com/axillusion/go-assigment/models"
	"github.com/axillusion/go-assigment/services"
)

type DataController struct {
	service services.DataServiceInterface
}

func NewDataController(service services.DataServiceInterface) *DataController {
	return &DataController{service: service}
}

// Endpoint to add a new dialog message
func (dataController *DataController) AddMessageToDialog(c *gin.Context) {
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
	go dataController.service.SaveToDB(customerID, dialogID, messageContent.Text, messageContent.Language, errorChannel)
	err = <-errorChannel
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to add to the database")
		return
	}
	c.JSON(http.StatusOK, "Message saved")
}

// Endpoint to fetch all the data
func (dataController *DataController) GetUserData(c *gin.Context) {
	language := c.Query("language")
	customerID := c.Query("customerID")
	var dataPoints []map[string]interface{}

	// Get the page number and page size from the query parameters
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	commons.Log.WithFields(logrus.Fields{"language": language, "customerID": customerID, "page": page, "pageSize": pageSize}).Info("Received request to fetch data on data/ with the following parameters")
	// Execute the query and scan the results into the dataPoints slice
	errorChannel := make(chan error)
	go dataController.service.FetchData(language, customerID, page, pageSize, &dataPoints, errorChannel)
	err := <-errorChannel
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	c.IndentedJSON(http.StatusOK, dataPoints)
}

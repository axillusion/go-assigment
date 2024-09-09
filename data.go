package main

import (
	"strconv"

	// Import to handle HTTP messages
	"net/http"

	//Import to extract the current time of the received messages

	// Import for using api endpoints
	"github.com/gin-gonic/gin"

	//Import for logging
	"github.com/sirupsen/logrus"

	//Import for database interaction
	"gorm.io/gorm"
)

// Basic struct containing the data stored for each user
type Data struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

type DialogRow struct {
	gorm.Model
	DialogID   string `json:"dialogId" gorm:"column:dialogID"`
	CustomerID string `json:"customerId" gorm:"column:customerID"`
	Text       string `json:"text" gorm:"column:stext"`
	Language   string `json:"language" gorm:"column:language"`
}

// Endpoint to fetch all the data
func GetUserData(c *gin.Context) {
	language := c.Query("language")
	customerID := c.Query("customerID")
	var dataPoints []map[string]interface{}

	query := db.Table("Dialog_rows").
		Select("dialogID", "customerID", "stext", "language").
		Order("created_at DESC")

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

	Log.WithFields(logrus.Fields{"language": language, "customerID": customerID, "page": page, "pageSize": pageSize}).Info("Received request to fetch data on data/ with the following parameters")

	db = db.Limit(pageSize).Offset(offset)
	// Execute the query and scan the results into the dataPoints slice
	if err := query.Find(&dataPoints).Error; err != nil {
		Log.Fatal("Failed to execute database query to extract the data")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.IndentedJSON(http.StatusOK, dataPoints)
}

// Deletes all the entries with the dialogID mentioned
func DeleteDialogData(dialogID string) {
	// query, err := db.Prepare("DELETE FROM Dialogs WHERE dialogID = ?")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer query.Close()
	// _, err = query.Exec(dialogID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	Log.Info("Deleting data from the database")
	res := db.Where("dialogID = ?", dialogID).Delete(&DialogRow{})
	if res.Error != nil {
		Log.WithFields(logrus.Fields{"dialogID": dialogID}).Fatal("Failed to delete rows from database with this id")
	}
}

// Adds a new dialog message
func AddMessageToDialog(c *gin.Context) {
	customerID := c.Param("customerID")
	dialogID := c.Param("dialogID")
	Log.WithFields(logrus.Fields{"customerID": customerID, "dialogID": dialogID}).Info("Post request for a new dialog entry with this params has been called")
	var messageContent Data
	err := c.BindJSON(&messageContent)
	// Message has to contain the text and the language of the message
	if err != nil {
		// Returns a bad request if the body is not valid
		c.JSON(http.StatusBadRequest, "Invalid JSON payload")
		Log.Warn("The JSON payload of the request invalid, processing finished")
		return
	}
	// save the new message received in the Database
	Log.Info("Processing request")
	defer SaveToDB(customerID, dialogID, messageContent.Text, messageContent.Language)
	c.JSON(http.StatusOK, "Message saved")
}

// Given a dialog entry, saves it in the database
func SaveToDB(customerID string, dialogID string, text string, language string) {
	// query, err := db.Prepare("INSERT INTO Dialogs (dialogID, customerID, stext, language, date_added) VALUES (?, ?, ?, ?, ?)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer query.Close()

	// _, err = query.Exec(dialogID, customerID, text, language, time.Now().Format("2006-01-02 15:04:05"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	Log.WithFields(logrus.Fields{"DialogID": dialogID, "CustomerID": "customerID", "Text": text, "Language": language}).Info("Creating new database entry with this information")
	err := db.Create(&DialogRow{DialogID: dialogID, CustomerID: customerID, Text: text, Language: language})
	db.Commit()
	if err.Error != nil {
		Log.Fatal("Failed to add the data to the database")
		return
	}
	Log.Info("Sucessfuly added the entry")
}

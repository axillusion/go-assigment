package main

import (
	"fmt"
	"log"
	"strconv"

	// Import to handle HTTP messages
	"net/http"

	//Import to extract the current time of the received messages

	// Import for using api endpoints
	"github.com/gin-gonic/gin"
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
	db = db.Limit(pageSize).Offset(offset)
	// Execute the query and scan the results into the dataPoints slice
	if err := query.Find(&dataPoints).Error; err != nil {
		fmt.Println("Failed to execute query:", err)
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
	res := db.Where("dialogID = ?", dialogID).Delete(&DialogRow{})
	if res.Error != nil {
		panic(res.Error)
	}
}

// Adds a new dialog message
func AddMessageToDialog(c *gin.Context) {
	customerID := c.Param("customerID")
	dialogID := c.Param("dialogID")
	var messageContent Data
	err := c.BindJSON(&messageContent)
	// Message has to contain the text and the language of the message
	if err != nil {
		// Returns a bad request if the body is not valid
		c.JSON(http.StatusBadRequest, "Invalid JSON payload")
		panic(err.Error())
		return
	}
	fmt.Println("Received message:", messageContent)
	// save the new message received in the Database
	SaveToDB(customerID, dialogID, messageContent.Text, messageContent.Language)
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
	err := db.Create(&DialogRow{DialogID: dialogID, CustomerID: customerID, Text: text, Language: language})
	db.Commit()
	if err.Error != nil {
		log.Fatal(err.RowsAffected)
	}
}

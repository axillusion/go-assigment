package services

import (
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"

	"main/commons"
	"main/models"
)

// Given a database with a search query on it, retrives the data
func FetchData(db *gorm.DB, destination interface{}, errorChannel chan<- error) {
	if err := db.Find(destination).Error; err != nil {
		commons.Log.Warn("Failed to execute database query to extract the data")
		errorChannel <- err
	}
}

// Deletes all the entries with the dialogID mentioned
func DeleteDialogData(dialogID string, errorChannel chan<- error) {
	commons.Log.Info("Deleting data from the database")
	res := commons.Db.Where("dialogID = ?", dialogID).Delete(&models.DialogRow{})
	if res.Error != nil {
		commons.Log.WithFields(logrus.Fields{"dialogID": dialogID}).Fatal("Failed to delete rows from database with this id")
		errorChannel <- res.Error
	}
}

// Given a dialog entry, saves it in the database
func SaveToDB(customerID string, dialogID string, text string, language string, errorChannel chan<- error) {
	commons.Log.WithFields(logrus.Fields{"DialogID": dialogID, "CustomerID": "customerID", "Text": text, "Language": language}).Info("Creating new database entry with this information")
	err := commons.Db.Create(&models.DialogRow{DialogID: dialogID, CustomerID: customerID, Text: text, Language: language})
	if err.Error != nil {
		commons.Log.WithFields(logrus.Fields{"error": err.Error}).Warn("Failed to add the data to the database because")
		errorChannel <- err.Error
		return
	}
	commons.Log.Info("Sucessfuly added the entry")
}

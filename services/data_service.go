package services

import (
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"

	"main/commons"
	"main/models"
)

// Given a database with a search query on it, retrives the data
// destination interface has to match the fetched data
func FetchData(db *gorm.DB, destination interface{}, errorChannel chan<- error) {
	commons.Log.Info("Fetching data")
	err := db.Find(destination).Error
	errorChannel <- err
	if err != nil {
		commons.Log.Warn("Failed to execute database query to extract the data")
		return
	}
	commons.Log.Info("Succesfully fetched the requested data")
}

// Deletes all the entries with the dialogID mentioned
func DeleteDialogData(dialogID string, errorChannel chan<- error) {
	commons.Log.Info("Deleting data from the database")
	res := commons.Db.Where("dialogID = ?", dialogID).Delete(&models.DialogRow{})
	errorChannel <- res.Error
	if res.Error != nil {
		commons.Log.WithFields(logrus.Fields{"dialogID": dialogID}).Fatal("Failed to delete rows from database with this id")
		return
	}
	commons.Log.Info("Successfully deleted from the database")
}

// Given a dialog entry, saves it in the database
func SaveToDB(customerID string, dialogID string, text string, language string, errorChannel chan<- error) {
	commons.Log.WithFields(logrus.Fields{"DialogID": dialogID, "CustomerID": "customerID", "Text": text, "Language": language}).Info("Creating new database entry with this information")
	err := commons.Db.Create(&models.DialogRow{DialogID: dialogID, CustomerID: customerID, Text: text, Language: language, Consent: false})
	errorChannel <- err.Error
	if err.Error != nil {
		commons.Log.WithFields(logrus.Fields{"error": err.Error}).Warn("Failed to add the data to the database because")
		return
	}
	commons.Log.Info("Sucessfuly added the entry")
}

// Modifies the consent of the user
func ModifyDB(dialogID string, errorChannel chan<- error) {
	commons.Log.WithFields(logrus.Fields{"DialogID": dialogID}).Info("Adding consent to this dialog")
	result := commons.Db.Table("dialog_rows").Where("dialogID = ?", dialogID).Update("consent", true)
	errorChannel <- result.Error
	if result.Error != nil {
		commons.Log.WithFields(logrus.Fields{"error": result.Error}).Warn("Failed to add the data to the database because")
		return
	}
	commons.Log.Info("Consent added succesfully")
}

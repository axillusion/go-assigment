package services

import (
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"

	"main/commons"
	"main/models"
)

type DataService struct {
	DB *gorm.DB
}

type DataServiceInterface interface {
	FetchData(language string, customerID string, page int, pageSize int, destination interface{}, errorChannel chan<- error)
	Delete(dialogID string, errorChannel chan<- error)
	SaveToDB(customerID string, dialogID string, text string, language string, errorChannel chan<- error)
	ModifyDB(dialogID string, errorChannel chan<- error)
}

// Constructor for the service to instantiate the database
func NewDataService(db *gorm.DB) *DataService {
	return &DataService{
		DB: db,
	}
}

// Given a database with a search query on it, retrives the data
// destination interface has to match the fetched data
func (dataService *DataService) FetchData(language string, customerID string, page int, pageSize int, destination interface{}, errorChannel chan<- error) {
	query := dataService.DB.Table("dialog_rows").
		Select("dialogID", "customerID", "stext", "language").
		Order("created_at DESC").
		Where("consent = ?", true)

	if language != "" {
		query.Where("language = ?", language)
	}

	if customerID != "" {
		query.Where("customerID = ?", customerID)
	}

	if pageSize > 0 {
		offset := pageSize * (page - 1)
		query.Limit(pageSize).Offset(offset)
	}

	commons.Log.Info("Fetching data")
	err := query.Find(destination).Error
	errorChannel <- err
	if err != nil {
		commons.Log.Warn("Failed to execute database query to extract the data")
		return
	}
	commons.Log.Info("Succesfully fetched the requested data")
}

// Deletes all the entries with the dialogID mentioned
func (dataService *DataService) Delete(dialogID string, errorChannel chan<- error) {
	commons.Log.Info("Deleting data from the database")
	res := dataService.DB.Where("dialogID = ?", dialogID).Delete(&models.DialogRow{})
	errorChannel <- res.Error
	if res.Error != nil {
		commons.Log.WithFields(logrus.Fields{"dialogID": dialogID}).Fatal("Failed to delete rows from database with this id")
		return
	}
	commons.Log.Info("Successfully deleted from the database")
}

// Given a dialog entry, saves it in the database
func (dataService *DataService) SaveToDB(customerID string, dialogID string, text string, language string, errorChannel chan<- error) {
	commons.Log.WithFields(logrus.Fields{"DialogID": dialogID, "CustomerID": "customerID", "Text": text, "Language": language}).Info("Creating new database entry with this information")
	err := dataService.DB.Create(&models.DialogRow{DialogID: dialogID, CustomerID: customerID, Text: text, Language: language, Consent: false})
	errorChannel <- err.Error
	if err.Error != nil {
		commons.Log.WithFields(logrus.Fields{"error": err.Error}).Warn("Failed to add the data to the database because")
		return
	}
	commons.Log.Info("Sucessfuly added the entry")
}

// Modifies the consent of the user
func (dataService *DataService) ModifyDB(dialogID string, errorChannel chan<- error) {
	commons.Log.WithFields(logrus.Fields{"DialogID": dialogID}).Info("Adding consent to this dialog")
	result := dataService.DB.Table("dialog_rows").Where("dialogID = ?", dialogID).Update("consent", true)
	errorChannel <- result.Error
	if result.Error != nil {
		commons.Log.WithFields(logrus.Fields{"error": result.Error}).Warn("Failed to add the data to the database because")
		return
	}
	commons.Log.Info("Consent added succesfully")
}

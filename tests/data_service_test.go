package tests

import (
	"fmt"
	"testing"

	"github.com/axillusion/go-assigment/models"
	"github.com/axillusion/go-assigment/services"
	"github.com/axillusion/go-assigment/utils"
	"gorm.io/gorm"
)

//var TestDialogData = models.DialogRow{CustomerID: "404", DialogID: "404", Text: "Test text", Language: "EN"}

func setupTestDB(t *testing.T) *services.DataService {
	_, service := utils.InitServerTestNoMock()
	err := service.DB.Create(&TestDialogData).Error
	if err != nil {
		t.Fatalf("Failed to set up test data: %v", err)
	}
	return service
}

func teardownTestDB(db *gorm.DB, t *testing.T) {
	err := db.Exec("DELETE FROM dialog_rows WHERE dialogID = ?", TestDialogData.DialogID).Error
	if err != nil {
		t.Fatalf("Failed to tear down test data: %v", err)
	}
}

func TestSaveToDB(t *testing.T) {
	service := setupTestDB(t)
	defer teardownTestDB(service.DB, t)

	errorChannel := make(chan error)
	go service.SaveToDB(TestDialogData.CustomerID, TestDialogData.DialogID, TestDialogData.Text, TestDialogData.Language, errorChannel)
	_ = <-errorChannel

	var entry models.DialogRow
	query := service.DB.Select("dialogID", "customerID", "stext", "language").Where("dialogID = ?", TestDialogData.DialogID)
	query.First(&entry)
	if entry.DialogID != TestDialogData.DialogID {
		t.Errorf("Failed to save to the database " + entry.DialogID)
	}
}

func TestDeleteDialogData(t *testing.T) {
	service := setupTestDB(t)
	defer teardownTestDB(service.DB, t)

	errorChannel := make(chan error)
	go service.Delete(TestDialogData.DialogID, errorChannel)
	_ = <-errorChannel

	var entry models.DialogRow
	query := service.DB.Select("dialogID", "customerID", "stext", "language").Where("dialogID = ?", TestDialogData.DialogID)
	query.First(&entry)
	if entry.DialogID == TestDialogData.DialogID {
		t.Errorf("Failed to delete from the database")
	}
}

func TestModifyDB(t *testing.T) {
	service := setupTestDB(t)
	defer teardownTestDB(service.DB, t)

	errorChannel := make(chan error)
	go service.ModifyDB(TestDialogData.DialogID, errorChannel)
	_ = <-errorChannel

	var entry models.DialogRow
	query := service.DB.Select("dialogID", "customerID", "stext", "language", "consent").Where("dialogID = ?", TestDialogData.DialogID)
	query.First(&entry)
	if entry.Consent != true {
		t.Errorf("Failed to modify the database entry")
	}
}

func TestFetchDataWithConsent(t *testing.T) {
	service := setupTestDB(t)
	defer teardownTestDB(service.DB, t)

	errorChannel := make(chan error)
	var data []models.DialogRow
	service.DB.Table("dialog_rows").Where("dialogID = ?", TestDialogData.DialogID).Update("consent", true)
	go service.FetchData(TestDialogData.Language, TestDialogData.CustomerID, 1, 10, &data, errorChannel)
	_ = <-errorChannel
	if len(data) == 0 {
		t.Errorf("Failed to fetch data from the database")
	}
}

func TestFetchDataWithoutConsent(t *testing.T) {
	service := setupTestDB(t)
	defer teardownTestDB(service.DB, t)

	errorChannel := make(chan error)
	var data []models.DialogRow
	go service.FetchData(TestDialogData.Language, TestDialogData.CustomerID, 1, 10, &data, errorChannel)
	_ = <-errorChannel
	if len(data) != 0 {
		t.Errorf("Failed to fetch data from the database")
	}
}

func TestFetchDataWithPagination(t *testing.T) {
	service := setupTestDB(t)
	defer teardownTestDB(service.DB, t)

	errorChannel := make(chan error)
	var data []models.DialogRow
	service.DB.Table("dialog_rows").Where("dialogID = ?", TestDialogData.DialogID).Update("consent", true)
	go service.FetchData(TestDialogData.Language, TestDialogData.CustomerID, 1, 1, &data, errorChannel)
	_ = <-errorChannel
	if len(data) != 1 {
		fmt.Print(len(data))
		t.Errorf("Failed to fetch data with pagination from the database")
	}
}

func TestFetchDataWithDifferentLanguage(t *testing.T) {
	service := setupTestDB(t)
	defer teardownTestDB(service.DB, t)

	errorChannel := make(chan error)
	var data []models.DialogRow
	service.DB.Table("dialog_rows").Where("dialogID = ?", TestDialogData.DialogID).Update("consent", true)
	go service.FetchData("FR", TestDialogData.CustomerID, 1, 10, &data, errorChannel)
	_ = <-errorChannel
	if len(data) != 0 {
		t.Errorf("Fetched data with incorrect language from the database")
	}
}

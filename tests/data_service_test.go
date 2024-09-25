package tests

import (
	"main/models"
	"main/utils"
	"testing"
)

func TestSaveToDB(t *testing.T) {
	_, service := utils.InitServerTestNoMock()
	errorChannel := make(chan error)
	go service.SaveToDB(TestDialogData.CustomerID, TestDialogData.DialogID, TestDialogData.Text, TestDialogData.Language, errorChannel)
	_ = <-errorChannel
	query := service.DB.Select("dialogID", "customerID", "stext", "language").Where("dialogID = ?", TestDialogData.DialogID)
	var entry models.DialogRow
	query.First(&entry)
	if entry.DialogID != TestDialogData.DialogID {
		t.Errorf("Failed to save to the database " + entry.DialogID)
	}
}

func TestDeleteDialogData(t *testing.T) {
	_, service := SetupTestNoMock()
	errorChannel := make(chan error)
	go service.Delete(TestDialogData.DialogID, errorChannel)
	_ = <-errorChannel
	query := service.DB.Select("dialogID", "customerID", "stext", "language").Where("dialogID = ?", TestDialogData.DialogID)
	var entry models.DialogRow
	query.First(&entry)
	if entry.DialogID == TestDialogData.DialogID {
		t.Errorf("Failed to delete from the database")
	}
}

func TestModifyDB(t *testing.T) {
	_, service := SetupTestNoMock()
	errorChannel := make(chan error)
	go service.ModifyDB(TestDialogData.DialogID, errorChannel)
	_ = <-errorChannel
	query := service.DB.Select("dialogID", "customerID", "stext", "language", "consent").Where("dialogID = ?", TestDialogData.DialogID)
	var entry models.DialogRow
	query.First(&entry)
	if entry.Consent != true {
		t.Errorf("Failed to modify the database entry")
	}
}

func TestFetchData(t *testing.T) {
	_, service := SetupTestNoMock()
	errorChannel := make(chan error)
	var data []models.DialogRow
	go service.FetchData(TestDialogData.Language, TestDialogData.CustomerID, 1, 10, &data, errorChannel)
	_ = <-errorChannel
	if len(data) == 0 {
		t.Errorf("Failed to fetch data from the database")
	}
}

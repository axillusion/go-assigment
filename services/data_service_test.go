package services

import (
	"testing"

	"github.com/axillusion/go-assigment/commons"
	"github.com/axillusion/go-assigment/logger"
	"github.com/axillusion/go-assigment/mocks"
	"github.com/axillusion/go-assigment/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

var TestDialogData = models.DialogRow{CustomerID: "404", DialogID: "404", Text: "Test text", Language: "EN"}
var router *gin.Engine
var service *DataService

func setupService() *DataService {
	commons.Log = logger.NewLogger()
	db := new(mocks.MockDatabase)
	service := NewDataService(db)
	db.On("Create", mock.Anything).Return(nil)
	db.On("Delete", mock.Anything, mock.Anything).Return(nil)
	db.On("Update", mock.Anything).Return(nil)
	db.On("Select", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	db.On("Where", mock.Anything, mock.Anything).Return(db)
	db.On("Table", mock.Anything).Return(db)
	db.On("Limit", mock.Anything).Return(db)
	db.On("Offset", mock.Anything).Return(db)
	db.On("Find", mock.Anything).Return(nil)
	db.On("Update", mock.Anything, mock.Anything).Return(nil)
	db.On("Order", mock.Anything).Return(db)
	return service
}

func TestSaveToDB(t *testing.T) {
	service := setupService()

	errorChannel := make(chan error)
	go service.SaveToDB(TestDialogData.CustomerID, TestDialogData.DialogID, TestDialogData.Text, TestDialogData.Language, errorChannel)
	err := <-errorChannel
	if err != nil {
		t.Errorf("SaveToDB failed")
	}
}

func TestDeleteDialogData(t *testing.T) {
	service := setupService()

	errorChannel := make(chan error)
	go service.Delete(TestDialogData.DialogID, errorChannel)
	err := <-errorChannel

	if err != nil {
		t.Errorf("Delete failed")
	}
}

func TestModifyDB(t *testing.T) {
	service := setupService()

	errorChannel := make(chan error)
	go service.ModifyDB(TestDialogData.DialogID, errorChannel)
	err := <-errorChannel

	if err != nil {
		t.Errorf("ModifyDB failed")
	}
}

func TestFetchDataWithConsent(t *testing.T) {
	service := setupService()

	errorChannel := make(chan error)
	var data []models.DialogRow
	service.DB.Table("dialog_rows").Where("dialogID = ?", TestDialogData.DialogID).Update("consent", true)
	go service.FetchData(TestDialogData.Language, TestDialogData.CustomerID, 1, 10, &data, errorChannel)
	err := <-errorChannel

	if err != nil {
		t.Errorf("FetchData with consent failed")
	}
}

func TestFetchDataWithoutConsent(t *testing.T) {
	service := setupService()

	errorChannel := make(chan error)
	var data []models.DialogRow
	go service.FetchData(TestDialogData.Language, TestDialogData.CustomerID, 1, 10, &data, errorChannel)
	err := <-errorChannel

	if err != nil {
		t.Errorf("FetchData without consent failed")
	}
}

func TestFetchDataWithPagination(t *testing.T) {
	service := setupService()

	errorChannel := make(chan error)
	var data []models.DialogRow
	service.DB.Table("dialog_rows").Where("dialogID = ?", TestDialogData.DialogID).Update("consent", true)
	go service.FetchData(TestDialogData.Language, TestDialogData.CustomerID, 1, 1, &data, errorChannel)
	err := <-errorChannel

	if err != nil {
		t.Errorf("FetchData with pagination failed")
	}
}

func TestFetchDataWithDifferentLanguage(t *testing.T) {
	service := setupService()

	errorChannel := make(chan error)
	var data []models.DialogRow
	service.DB.Table("dialog_rows").Where("dialogID = ?", TestDialogData.DialogID).Update("consent", true)
	go service.FetchData("FR", TestDialogData.CustomerID, 1, 10, &data, errorChannel)
	err := <-errorChannel

	if err != nil {
		t.Errorf("FetchData with different language failed")
	}
}

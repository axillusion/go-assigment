package tests

import (
	"github.com/axillusion/go-assigment/models"
	"github.com/axillusion/go-assigment/services"
	"github.com/axillusion/go-assigment/utils"

	"github.com/gin-gonic/gin"
)

var TestDialogData = models.DialogRow{CustomerID: "404", DialogID: "404", Text: "Test text", Language: "EN"}

func SetupTestNoMock() (*gin.Engine, *services.DataService) {
	routes, service := utils.InitServerTestNoMock()
	service.DB.Create(&TestDialogData)
	return routes, service
}

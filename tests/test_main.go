package tests

import (
	"main/models"
	"main/services"
	"main/utils"

	"github.com/gin-gonic/gin"
)

var TestDialogData = models.DialogRow{CustomerID: "404", DialogID: "404", Text: "Test text", Language: "EN"}

func SetupTestNoMock() (*gin.Engine, *services.DataService) {
	routes, service := utils.InitServerTestNoMock()
	service.DB.Create(&TestDialogData)
	return routes, service
}

package routes

import (
	"main/controllers"
	"main/services"

	"github.com/gin-gonic/gin"
)

func InitEndpoints(service services.DataServiceInterface) *gin.Engine {
	dataController := controllers.NewDataController(service)
	consentController := controllers.NewConsentController(service)
	router := gin.Default()
	router.GET("/data", dataController.GetUserData)
	router.POST("/consents/:dialogID", consentController.CheckConsent)
	router.POST("/data/:customerID/:dialogID", dataController.AddMessageToDialog)
	//router.Run("localhost:8080")
	return router
}

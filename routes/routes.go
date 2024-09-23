package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func InitEndpoints() {
	router := gin.Default()
	router.GET("/data", controllers.GetUserData)
	router.POST("/consents/:dialogID", controllers.CheckConsent)
	router.POST("/data/:customerID/:dialogID", controllers.AddMessageToDialog)
	router.Run("localhost:8080")
}

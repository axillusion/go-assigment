package main

import (
	"github.com/axillusion/go-assigment/commons"
	"github.com/axillusion/go-assigment/controllers"
	"github.com/axillusion/go-assigment/database"
	"github.com/axillusion/go-assigment/logger"
	"github.com/axillusion/go-assigment/services"
	"github.com/gin-gonic/gin"
)

func main() {
	commons.Log = logger.NewLogger()
	db := new(database.GormDatabase)
	db.Connect()
	defer db.Close()
	service := services.NewDataService(db)
	dataController := controllers.NewDataController(service)
	consentController := controllers.NewConsentController(service)
	router := gin.Default()
	router.GET("/data", dataController.GetUserData)
	router.POST("/consents/:dialogID", consentController.CheckConsent)
	router.POST("/data/:customerID/:dialogID", dataController.AddMessageToDialog)
	router.Run(":8080")
}

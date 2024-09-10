package main

import (

	// Import for using api endpoints
	"github.com/gin-gonic/gin"

	//Import for logging

	//Imports for database interaction

	"go-assigment/consents"
	"go-assigment/data"
)

func main() {
	data.Init()

	router := gin.Default()
	router.GET("/data", data.GetUserData)
	router.POST("/consents/:dialogID", consents.CheckConsent)
	router.POST("/data/:customerID/:dialogID", data.AddMessageToDialog)
	router.Run("localhost:8080")
}

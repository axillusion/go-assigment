package main

import (
	"os"
	// Import for using api endpoints
	"github.com/gin-gonic/gin"

	//Import for logging
	"github.com/sirupsen/logrus"

	//Imports for database interaction
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-assigment/consents"
	"go-assigment/data"
)

func main() {
	data.Log = logrus.New()
	data.Log.SetOutput(os.Stdout)
	data.Log.SetFormatter(&logrus.JSONFormatter{})
	data.Log.SetLevel(logrus.WarnLevel)

	var err error
	data.Db, err = gorm.Open(mysql.Open("root:sqlTudelft+990!@tcp(localhost:3306)/Messages"), &gorm.Config{})
	data.Db.AutoMigrate(&data.DialogRow{})
	if err != nil {
		data.Log.Fatal("Could not connect to the database")
	}
	data.Log.Info("Succesfully established database connection")

	router := gin.Default()
	router.GET("/data", data.GetUserData)
	router.POST("/consents/:dialogID", consents.CheckConsent)
	router.POST("/data/:customerID/:dialogID", data.AddMessageToDialog)
	router.Run("localhost:8080")
}

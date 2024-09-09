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
)

var db *gorm.DB
var Log *logrus.Logger

func main() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.WarnLevel)

	var err error
	db, err = gorm.Open(mysql.Open("root:sqlTudelft+990!@tcp(localhost:3306)/Messages"), &gorm.Config{})
	db.AutoMigrate(&DialogRow{})
	if err != nil {
		Log.Fatal("Could not connect to the database")
	}
	Log.Info("Succesfully established database connection")

	router := gin.Default()
	router.GET("/data", GetUserData)
	router.POST("/consents/:dialogID", CheckConsent)
	router.POST("/data/:customerID/:dialogID", AddMessageToDialog)
	router.Run("localhost:8080")
}

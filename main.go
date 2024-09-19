package main

import (
	"os"
	// Import for using api endpoints
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	//Import for logging
	"github.com/sirupsen/logrus"

	//Imports for database interaction
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"main/consents"
	"main/data"
)

func main() {
	data.Log = logrus.New()
	data.Log.SetOutput(os.Stdout)
	data.Log.SetFormatter(&logrus.JSONFormatter{})
	data.Log.SetLevel(logrus.WarnLevel)

	var err error

	err = godotenv.Load()
	if err != nil {
		data.Log.Fatal("Could not load env file")
	}

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_string := db_user + ":" + db_pass + "@tcp(" + db_port + ")/" + db_name
	data.Db, err = gorm.Open(mysql.Open(db_string), &gorm.Config{})
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

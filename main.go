package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	var err error
	db, err = gorm.Open(mysql.Open("root:sqlTudelft+990!@tcp(localhost:3306)/Messages"), &gorm.Config{})
	db.AutoMigrate(&DialogRow{})
	if err != nil {
		panic(err.Error())
	}

	router := gin.Default()
	router.GET("/data", GetUserData)
	router.POST("/consents/:dialogID", CheckConsent)
	router.POST("/data/:customerID/:dialogID", AddMessageToDialog)
	router.Run("localhost:8080")
}

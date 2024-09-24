package utils

import (
	"main/commons"
	"main/models"
	"main/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitServer() {
	initLogger()
	initDB()
	routes.InitEndpoints()
}

func initDB() {
	var err error

	err = godotenv.Load()
	if err != nil {
		commons.Log.Fatal("Could not load env file")
	}

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_string := db_user + ":" + db_pass + "@tcp(" + db_port + ")/" + db_name
	commons.Db, err = gorm.Open(mysql.Open(db_string), &gorm.Config{})
	commons.Db.AutoMigrate(&models.DialogRow{})
	if err != nil {
		commons.Log.Fatal("Could not connect to the database")
	}
	commons.Log.Info("Succesfully established database connection")
}

func initLogger() {
	commons.Log = logrus.New()
	commons.Log.SetOutput(os.Stdout)
	commons.Log.SetFormatter(&logrus.JSONFormatter{})
	commons.Log.SetLevel(logrus.DebugLevel)
}

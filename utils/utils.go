package utils

import (
	"os"
	"path/filepath"

	"github.com/axillusion/go-assigment/commons"
	"github.com/axillusion/go-assigment/models"
	"github.com/axillusion/go-assigment/routes"
	"github.com/axillusion/go-assigment/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func loadEnv() {
	envPath := ".env"
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		envPath = filepath.Join("..", ".env")
	}
	err := godotenv.Load(envPath)
	if err != nil {
		commons.Log.Fatalf("Could not load env file from path %s with error: %v", envPath, err)
	} else {
		commons.Log.Infof("Loaded env file from path %s", envPath)
	}
}

func initDB() *gorm.DB {
	var err error

	loadEnv()

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_string := db_user + ":" + db_pass + "@tcp(" + db_port + ")/" + db_name
	DB, err := gorm.Open(mysql.Open(db_string), &gorm.Config{})
	DB.AutoMigrate(&models.DialogRow{})
	if err != nil {
		commons.Log.Fatal("Could not connect to the database")
	}
	commons.Log.Info("Succesfully established database connection")
	return DB
}

func InitServer() {
	initLogger()
	DB := initDB()
	service := services.NewDataService(DB)
	routes.InitEndpoints(service)
}

func InitServerTestWithMock() (*gin.Engine, *commons.MockDataService) {
	//fmt.Println(os.Getwd())
	initLogger()
	initDB()
	service := new(commons.MockDataService)
	return routes.InitEndpointsTest(service), service
}

func InitServerTestNoMock() (*gin.Engine, *services.DataService) {
	//fmt.Println(os.Getwd())
	initLogger()
	DB := initDB()
	service := services.NewDataService(DB)
	return routes.InitEndpointsTest(service), service
}

func initLogger() {
	commons.Log = logrus.New()
	commons.Log.SetOutput(os.Stdout)
	commons.Log.SetFormatter(&logrus.JSONFormatter{})
	commons.Log.SetLevel(logrus.DebugLevel)
}

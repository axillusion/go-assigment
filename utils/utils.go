package utils

import (
	"main/commons"
	"main/models"
	"main/routes"
	"main/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitServer() {
	initLogger()
	DB := initDB()
	service := services.NewDataService(DB)
	routes.InitEndpoints(service)
}

func InitServerTestWithMock() (*gin.Engine, *commons.MockDataService) {
	os.Chdir("..")
	//fmt.Println(os.Getwd())
	initLogger()
	initDB()
	service := new(commons.MockDataService)
	return routes.InitEndpointsTest(service), service
}

func InitServerTestNoMock() (*gin.Engine, *services.DataService) {
	os.Chdir("..")
	//fmt.Println(os.Getwd())
	initLogger()
	DB := initDB()
	service := services.NewDataService(DB)
	return routes.InitEndpointsTest(service), service
}

func initDB() *gorm.DB {
	var err error

	err = godotenv.Load()
	if err != nil {
		commons.Log.Fatal("Could not load env file with error ?", err)
	}

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

func initLogger() {
	commons.Log = logrus.New()
	commons.Log.SetOutput(os.Stdout)
	commons.Log.SetFormatter(&logrus.JSONFormatter{})
	commons.Log.SetLevel(logrus.DebugLevel)
}

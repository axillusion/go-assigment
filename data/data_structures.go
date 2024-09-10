package data

import (
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Basic struct containing the data stored for each user
type Data struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

type DialogRow struct {
	gorm.Model
	DialogID   string `json:"dialogId" gorm:"column:dialogID"`
	CustomerID string `json:"customerId" gorm:"column:customerID"`
	Text       string `json:"text" gorm:"column:stext"`
	Language   string `json:"language" gorm:"column:language"`
}

var Db *gorm.DB
var Log *logrus.Logger

func Init() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.WarnLevel)
	var err error
	Db, err = gorm.Open(mysql.Open("root:sqlTudelft+990!@tcp(localhost:3306)/Messages"), &gorm.Config{})
	Db.AutoMigrate(&DialogRow{})

	if err != nil {
		Log.Fatal("Could not connect to the database")
	}
	Log.Info("Succesfully established database connection")
}

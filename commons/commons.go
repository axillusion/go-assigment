package commons

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Log *logrus.Logger

package models

import (
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

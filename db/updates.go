package db

import "github.com/jinzhu/gorm"

// Update represents an update from Telegram
type Update struct {
	gorm.Model
	UpdateID int `gorm:"not null;unique"`
	From     int `gorm:"not null"`
	Chat     int `gorm:"not null"`
	Message  string
}

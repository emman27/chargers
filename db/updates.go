package db

import "github.com/jinzhu/gorm"

// Update represents an update from Telegram
type Update struct {
	gorm.Model
	UpdateID         int
	Message          string
	EditedMessage    string
	ShippingQuery    string
	PreCheckoutQuery string
}

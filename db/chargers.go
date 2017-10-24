package db

import "github.com/jinzhu/gorm"

// Charger table contains information about chargers in stock
type Charger struct {
	gorm.Model
	Brand       string
	ModelNumber string
	Voltage     string
	BelongsToID uint
	Type        string
}

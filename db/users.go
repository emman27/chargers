package db

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// User represents an user from Telegram
type User struct {
	gorm.Model
	UserID       int `gorm:"not null;unique_index"`
	IsBot        bool
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	Chargers     []Charger `gorm:"ForeignKey:BelongsToID"`
}

func (u User) String() string {
	return strings.Trim(fmt.Sprint(u.FirstName, " ", u.LastName), " ")
}

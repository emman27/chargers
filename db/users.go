package db

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

// User represents an user from Telegram
type User struct {
	gorm.Model
	UserID string
	IsBot bool
	FirstName string
	LastName string
	Username string
	LanguageCode string
}

func (u User) String() string {
	return fmt.Sprintf(u.FirstName, " ", u.LastName)
}
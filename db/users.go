package db

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"strings"
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
	return strings.Trim(fmt.Sprint(u.FirstName, " ", u.LastName), " ")
}
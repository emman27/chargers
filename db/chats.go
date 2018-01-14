package db

import "github.com/jinzhu/gorm"

// Chat reflects a chat. More meaningful than users cos of group chats
type Chat struct {
	gorm.Model
	ChatID int
	State  State
}

// State reflects the current status of a chat
type State int

// Valid states
const (
	NONE      State = 0
	BORROWING State = 1
	SHARING   State = 2
)

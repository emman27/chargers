package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/emman27/chargers/api"
	"github.com/emman27/chargers/db"
	"github.com/jinzhu/gorm"
)

// Receiver allows dependency injection for the database
type Receiver struct {
	DB *gorm.DB
}

// Receiver receives POST request from Telegram's callbacks
func (rcv *Receiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var update api.UpdateSchema
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Fatal("JSON decoding failed: ", err)
	}
	upd := db.Update{
		Message:  update.Message.Text,
		From:     update.Message.From.ID,
		Chat:     update.Message.Chat.ID,
		UpdateID: update.UpdateID,
	}
	rcv.DB.Create(&upd)

	var user db.User
	rcv.DB.First(&user, "user_id = ?", upd.From)
	if user.UserID == 0 {
		go log.Println("New user!")
		user = db.User{
			UserID:    update.Message.From.ID,
			FirstName: update.Message.From.FirstName,
			LastName:  update.Message.From.LastName,
		}
		rcv.DB.Create(&user)
	}
	go log.Println(user.String(), "sent a message:", upd.Message)

	go parseUpdate(&upd)
}

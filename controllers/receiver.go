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
	log.Println(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Fatal("JSON decoding failed: ", err)
	}
	obj := db.Update{
		Message:  update.Message.Text,
		From:     update.Message.From.ID,
		Chat:     update.Message.Chat.ID,
		UpdateID: update.UpdateID,
	}
	rcv.DB.Create(&obj)
	api.Reply(obj.Chat, "Thanks for sending us a message! You sent us: "+obj.Message)
}

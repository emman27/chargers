package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emman27/chargers/api"
	"github.com/emman27/chargers/constants"
	"github.com/emman27/chargers/db"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Receiver allows dependency injection for the database
type Receiver struct {
	DB *gorm.DB
}

// Receiver receives POST request from Telegram's callbacks
func (rcv *Receiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var update api.UpdateSchema
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		logrus.Error("JSON decoding failed: ", err, r.Body)
		return
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
		go logrus.Info("New user!")
		user = db.User{
			UserID:    update.Message.From.ID,
			FirstName: update.Message.From.FirstName,
			LastName:  update.Message.From.LastName,
		}
		rcv.DB.Create(&user)
	}
	logged := fmt.Sprintf("%s (%d) sent a message: %s", user.String(), update.Message.Chat.ID, upd.Message)
	go logrus.Info(logged)
	go api.Reply(constants.AdminChatID, logged)
	go (&Parser{rcv.DB}).parseUpdate(&upd)
}

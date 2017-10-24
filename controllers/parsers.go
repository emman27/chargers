package controllers

import (
	"fmt"
	"strings"

	"github.com/emman27/chargers/api"
	"github.com/emman27/chargers/db"
	"github.com/jinzhu/gorm"
)

// Parser implements gorm.DB
type Parser struct{ *gorm.DB }

func (p *Parser) parseUpdate(update *db.Update) {
	switch arr := strings.Fields(update.Message); arr[0] {
	case "/start", "/help":
		go api.Reply(update.Chat, "Welcome to the Chargers Bot!\nNice to meet you!\n\n- To share a charger, type /share <product>\n- To borrow a charger, type /borrow <product>\nFor example, you can type `/share android` to share a Android charger")
	case "/borrow":
		go api.Reply(update.Chat, "Heh this doesn't work yet")
	case "/share":
		go p.parseShare(update, strings.Join(arr[1:], " "))
	case "/stock":
		go api.Reply(update.Chat, "Hold on, lemme fetch you the stock")
	default:
		go api.Reply(update.Chat, "That command is not recognized.")
	}
}

func (p *Parser) parseShare(update *db.Update, chargerFor string) {
	charger := db.Charger{Type: chargerFor}
	go p.DB.Create(&charger)
	go api.Reply(update.Chat, fmt.Sprintf("Your charger for %s has now been listed!", chargerFor))
}

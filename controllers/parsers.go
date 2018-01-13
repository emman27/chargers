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
		go api.ReplyWithOptions(update.Chat, "Hello! Thank you for subscribing to Jaetaan! A platform for you to share you chargers and earn money or loan a charger for a small fee! Try it out now!", []string{
			"Share a charger",
			"Borrow a charger",
		})
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

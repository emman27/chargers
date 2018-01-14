package controllers

import (
	"fmt"

	"github.com/emman27/chargers/api"
	"github.com/emman27/chargers/db"
	"github.com/jinzhu/gorm"
)

// Parser implements gorm.DB
type Parser struct{ *gorm.DB }

func (p *Parser) parseUpdate(update *db.Update) {
	switch update.Message {
	case "/start", "/help":
		go api.ReplyWithOptions(update.Chat, "Hello! Thank you for subscribing to Jaetaan! A platform for you to share you chargers and earn money or loan a charger for a small fee! Try it out now!", []string{
			"Share a charger",
			"Borrow a charger",
		})
	case "/stock":
		go api.Reply(update.Chat, "Hold on, lemme fetch you the stock")
	case "Share a charger":
		go api.Reply(update.Chat, "Thanks for offering to share a charger! Please write us a short description on the charger you would like to share. Try to be as detailed as possible so we can match you with someone!")
	case "Borrow a charger":
		go api.Reply(update.Chat, "Could you give us a description of the charger you'd like to borrow? Try to be as specific as possible so that we can find you a good match!")
	default:
		go api.Reply(update.Chat, "That command is not recognized.")
	}
}

func (p *Parser) parseShare(update *db.Update, chargerFor string) {
	charger := db.Charger{Type: chargerFor}
	go p.DB.Create(&charger)
	go api.Reply(update.Chat, fmt.Sprintf("Your charger for %s has now been listed!", chargerFor))
}

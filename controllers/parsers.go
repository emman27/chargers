package controllers

import (
	"strings"

	"github.com/emman27/chargers/api"
	"github.com/emman27/chargers/db"
)

func parseUpdate(update *db.Update) {
	switch strings.Fields(update.Message)[0] {
	case "/start", "/help":
		go api.Reply(update.Chat, "Welcome to the Chargers Bot!\nNice to meet you!\n\n- To share a charger, type /share android\n- To borrow a charger, type /borrow iphone7")
	case "/borrow":
		go api.Reply(update.Chat, "Heh this doesn't actually work")
	case "/share":
		go api.Reply(update.Chat, "Heh this doesn't work either")
	case "/stock":
		go api.Reply(update.Chat, "Hold on, lemme fetch you the stock")
	default:
		go api.Reply(update.Chat, "That command is not recognized.")
	}
}

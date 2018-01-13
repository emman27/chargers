package api

import "github.com/parnurzeal/gorequest"
import "log"

// Reply returns a chat message
func Reply(chatID int, text string) {
	data := map[string]interface{}{
		"chat_id":    chatID,
		"text":       text,
		"parse_mode": "Markdown",
	}
	_, _, errs := gorequest.New().
		Post(baseURL+"sendMessage").
		Set("Content-Type", "application/json").
		Send(data).
		End()
	if len(errs) != 0 {
		log.Fatal("Error replying: ", errs)
	}
}

// ReplyWithOptions returns a chat message with an in-app keyboard
func ReplyWithOptions(chatID int, text string, options []string) {
	data := map[string]interface{}{
		"chat_id":    chatID,
		"text":       text,
		"parse_mode": "Markdown",
		"reply_markup": ReplyKeyboardMarkup{
			OneTimeKeyboard: true,
			Keyboard: [][]KeyboardButton{
				[]KeyboardButton{KeyboardButton{Text: "Share a charger"}},
				[]KeyboardButton{KeyboardButton{Text: "Borrow a charger"}},
			},
		},
	}
	_, _, errs := gorequest.New().
		Post(baseURL+"sendMessage").
		Set("Content-Type", "application/json").
		Send(data).
		End()
	if len(errs) != 0 {
		log.Fatal("Error replying: ", errs)
	}
}

package api

import (
	"encoding/json"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/parnurzeal/gorequest"
)

// GetUpdates pings the telegram server for the latest updates
func GetUpdates(ch chan []UpdateSchema) {
	const baseURL = "https://api.telegram.org/bot359390703:AAHbvNwIrh4M97IEvbhZb1ZvBDygNs50I20/getUpdates"

	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("Database did not open: ", err)
	}
	defer database.Close()

	offset := database.Exec("SELECT max(update_id) FROM updates;")
	log.Println("Current latest update: ", offset)

	_, body, errors := gorequest.New().Get(baseURL).End()
	var result updateResponse
	if len(errors) > 0 {
		log.Fatal("Request failed: ", errors)
	}
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatal("Failed to parse json: ", body, err)
	}
	ch <- result.Result
}

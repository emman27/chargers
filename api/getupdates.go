package api

import (
	"encoding/json"
	"fmt"
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

	type dbResult struct {
		Max      int
		UpdateID int
	}
	var queryResult dbResult
	database.Raw("SELECT MAX(update_id) as max, update_id FROM updates").Scan(&queryResult)
	log.Println("Current latest update: ", queryResult.Max)

	_, body, errors := gorequest.New().Get(baseURL).Query(fmt.Sprintf("offset=%d", queryResult.Max+1)).End()
	var result updateResponse
	if len(errors) > 0 {
		log.Fatal("Request failed: ", errors)
	}
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatal("Failed to parse json: ", body, err)
	}
	ch <- result.Result
}

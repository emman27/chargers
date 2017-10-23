package main

import (
	"log"

	"github.com/emman27/chargers/api"
	"github.com/emman27/chargers/constants"
	"github.com/emman27/chargers/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	database, err := gorm.Open(constants.Driver, constants.DBName)
	if err != nil {
		log.Fatal("Database did not open: ", err)
	}
	defer database.Close()
	ch := make(chan []api.UpdateSchema)
	go api.GetUpdates(ch)
	results := <-ch
	for _, result := range results {
		database.Create(&db.Update{UpdateID: result.UpdateID, From: result.Message.From.ID, Message: result.Message.Text})
	}
}

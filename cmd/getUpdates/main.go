package main

import (
	"fmt"
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
		log.Println(result)
		update := db.Update{
			UpdateID: result.UpdateID,
			From:     result.Message.From.ID,
			Message:  result.Message.Text,
			Chat:     result.Message.Chat.ID,
		}
		database.Create(&update)
		api.Reply(update.Chat, fmt.Sprint("Thanks for sending us a message! Your message: ", update.Message))
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/emman27/chargers/api"
	"github.com/emman27/chargers/controllers"
	"github.com/emman27/chargers/db"
)

func main() {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer database.Close()

	database.AutoMigrate(&db.Update{})
	database.AutoMigrate(&db.Charger{})
	database.AutoMigrate(&db.User{})
	database.AutoMigrate(&db.Chat{})

	http.Handle("/update", &controllers.Receiver{DB: database})
	api.SetWebhook("https://904d8327.ngrok.io/update")
	defer api.DeleteWebhook()
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("Server running!")
}

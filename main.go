package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/emman27/charger/db"
)

func main() {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer database.Close()

	database.AutoMigrate(&db.Update{})
	database.AutoMigrate(&db.Charger{})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

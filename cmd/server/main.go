package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gorilla/mux"

	"github.com/emman27/chargers/db"
	"github.com/emman27/chargers/controllers"
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

	router := mux.NewRouter()
	router.Handle("receiver", &controllers.Receiver{DB: database}).Methods("POST")

	server := &http.Server{
		Handler: router,
		Addr: "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15* time.Second,
	}
	log.Println("Server running!")

	log.Fatal(server.ListenAndServe())
}

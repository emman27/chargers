package main

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"

	"github.com/emman27/chargers/api"
	"github.com/emman27/chargers/controllers"
	"github.com/emman27/chargers/db"
)

func main() {
	database := setupDatabase()
	redis := setupRedis()
	logrus.Info(redis)
	http.Handle("/update", &controllers.Receiver{DB: database})
	api.SetWebhook("https://def2f71b.ngrok.io")
	logrus.Info("Server starting now")
	defer api.DeleteWebhook()
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

func setupDatabase() *gorm.DB {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer database.Close()

	database.AutoMigrate(&db.Update{})
	database.AutoMigrate(&db.Charger{})
	database.AutoMigrate(&db.User{})
	database.AutoMigrate(&db.Chat{})
	return database
}

func setupRedis() *redis.Client {
	cl := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	statusCmd := cl.Ping()
	if err := statusCmd.Err(); err != nil {
		panic(err)
	}
	return cl
}

package controllers

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

// Receiver allows dependency injection for the database
type Receiver struct {
	DB *gorm.DB
}

// Receiver receives POST request from Telegram's callbacks
func (rcv *Receiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
}
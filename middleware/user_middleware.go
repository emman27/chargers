package middleware

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

// Middleware to extract out the current user
type UserMiddleware struct {
	db *gorm.DB
}

func (m *UserMiddleware) ServeHTTP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		h.ServeHTTP(w, r)
	})
}
package db

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func TestUser_String(t *testing.T) {
	type fields struct {
		Model        gorm.Model
		UserID       string `gorm:"unique_index"`
		IsBot        bool
		FirstName    string
		LastName     string
		Username     string
		LanguageCode string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"First name only", fields{UserID: "1", FirstName: "Emmanuel"}, "Emmanuel"},
		{"Last name only", fields{UserID: "2", FirstName: "Goh"}, "Goh"},
		{"First and Last Names", fields{UserID: "3", FirstName: "Emmanuel", LastName: "Goh"}, "Emmanuel Goh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				Model:        tt.fields.Model,
				UserID:       tt.fields.UserID,
				IsBot:        tt.fields.IsBot,
				FirstName:    tt.fields.FirstName,
				LastName:     tt.fields.LastName,
				Username:     tt.fields.Username,
				LanguageCode: tt.fields.LanguageCode,
			}
			if got := u.String(); got != tt.want {
				t.Errorf("User.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

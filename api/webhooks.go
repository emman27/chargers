package api

import (
	"log"

	"github.com/parnurzeal/gorequest"
)

// SetWebhook sets the webhook to a given URL
func SetWebhook(url string) (bool, []error) {
	data := map[string]string{
		"url": url,
	}

	_, _, errs := gorequest.New().Post(baseURL+"setWebhook").
		Set("Content-Type", "application/json").
		Send(data).
		End()
	if len(errs) != 0 {
		log.Fatal("Create webhook failed: ", errs)
		return false, errs
	}
	log.Println("Webhook created: " + url)
	return true, make([]error, 0)
}

// DeleteWebhook removes the webhook from the system
func DeleteWebhook() (bool, []error) {
	url := baseURL + "deleteWebhook"
	_, _, err := gorequest.New().Post(url).End()
	if err != nil {
		log.Fatal("Delete webhook failed: ", err)
		return false, err
	}
	log.Println("Webhook deleted")
	return true, make([]error, 0)
}

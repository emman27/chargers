package api

import (
	"log"

	"github.com/parnurzeal/gorequest"
)

// SetWebhook sets the webhook to a given URL
func SetWebhook(url string, ch chan bool) {
	data := map[string]string{
		"url": url,
	}

	done := func(_ gorequest.Response, body string, errs []error) {
		if len(errs) != 0 {
			log.Fatal("Create webhook failed: ", errs)
		}
		log.Println(body)
	}

	gorequest.New().Post(baseURL+"setWebhook").
		Set("Content-Type", "application/json").
		Send(data).
		End(done)
}

// DeleteWebhook removes the webhook from the system
func DeleteWebhook(ch chan bool) {
	url := baseURL + "deleteWebhook"
	_, _, err := gorequest.New().Post(url).End()
	if err != nil {
		log.Fatal("Delete webhook failed: ", err)
		ch <- false
	}
	log.Println("Webhook deleted")
	ch <- true
}

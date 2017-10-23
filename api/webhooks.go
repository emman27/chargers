package api

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const baseURL = "https://api.telegram.org/359390703:AAHbvNwIrh4M97IEvbhZb1ZvBDygNs50I20/bot/"

// SetWebhook sets the webhook to a given URL
func SetWebhook(url string, ch chan bool) {
	data := map[string]string{
		"url": url,
	}

	done := func(resp gorequest.Response, body string, errs []error) {
		fmt.Println(resp.Status)
	}

	gorequest.New().Post(baseURL+"setWebhook").
		Set("Content-Type", "application/json").
		Send(data).
		End(done)
}

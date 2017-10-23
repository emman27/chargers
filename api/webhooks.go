package api

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// SetWebhook sets the webhook to a given URL
func SetWebhook(url string, ch chan bool) {
	baseURL := "https://api.telegram.org/bot359390703:AAHbvNwIrh4M97IEvbhZb1ZvBDygNs50I20/"
	
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

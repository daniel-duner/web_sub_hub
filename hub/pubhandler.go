package main

import (
	"bytes"
	"fmt"
	"net/http"
)

//Handles incoming publishing requests, incorrect requets generates a bad request
func publishingHandler(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "POST") || (r.Header.Get("Content-Type") != "application/x-www-form-urlencoded") {
		fmt.Fprintf(w, "Wrong Method, received method: "+r.Method)
		w.WriteHeader(400)
	} else {
		publish(r)
	}
}

//Sends a JSON message to to all clients subscribed to the topic, it creates and passes a HMAC signature in the request header
func publish(r *http.Request) {
	r.ParseForm()
	message := r.Form.Get("message")
	hubURL := r.Form.Get("hub.url")
	messageJSON := []byte(`{"message":"` + message + `"}`)

	for cb, sub := range subscriptions {
		if sub.topic == hubURL {
			signature := encryptHMAC(messageJSON, sub.secret)
			req, err := http.NewRequest("POST", cb, bytes.NewBuffer(messageJSON))
			req.Header.Set("X-Hub-Signature", "sha256="+signature)
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("client: "+cb, err)
				fmt.Println("client: ", resp.Status)
			}
			defer resp.Body.Close()
		}
	}
}

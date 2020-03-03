package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Handles incoming subscription requests
func subscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if err := validateRequest(r); err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(400)
	} else {
		go func(r *http.Request) {
			serviceHandler(r)
		}(r)
	}
}

//Directs requests to the correct service
func serviceHandler(r *http.Request) {
	r.ParseForm()
	cb := r.Form.Get("hub.callback")
	mode := r.Form.Get("hub.mode")
	topic := r.Form.Get("hub.topic")
	secret := r.Form.Get("hub.secret")
	if err := verificationRequest(cb, mode, topic, secret); err != nil {
		fmt.Println(err)
	} else {
		if mode == "subscribe" {
			sub := subscription{cb, mode, topic, secret}
			subscriptions[cb] = sub
		} else if mode == "unsubscribe" {
			delete(subscriptions, cb)
		}
		fmt.Println(mode + " Success")
	}
}

//Sends a request to an alleged subscriber to verify their intent, returns nil if all is well
func verificationRequest(cb string, mode string, topic string, secret string) error {
	challenge := randomString(128)
	verificationURL := cb + "?hub.mode=" + mode + "&hub.topic=" + topic + "&hub.secret=" + secret + "&hub.challenge=" + challenge
	resp, err := http.Get(verificationURL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode >= 200 || resp.StatusCode <= 299 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return err
		}
		chal := string(body)
		if chal != challenge {
			return errors.New("Challenge failed to pass received challenge: " + chal)
		}
		return nil
	}
	return errors.New("Verification of intent failed action " + mode + " did not go through")
}

//Validates correct format of a subscribtion request, returns nil if all is well
func validateRequest(r *http.Request) error {
	h := r.Header
	r.ParseForm()
	params := r.Form
	cb := params.Get("hub.callback")
	mode := params.Get("hub.mode")
	topic := params.Get("hub.topic")
	if r.Method != "POST" {
		fmt.Println("Error: Received a " + r.Method + " request from" + h.Get("X-FORWARDED-FOR"))
		return errors.New("Wrong request method, received method: " + r.Method)
	} else if h.Get("Content-Type") != "application/x-www-form-urlencoded" {
		fmt.Print("Error: received Content-Type: " + h.Get("Content-Type"))
		return errors.New("Wrong Content-Type, received was type: " + h.Get("Content-Type"))
	} else if (cb == "") || (mode == "") || (topic == "") {
		fmt.Print("Error: missing parameter")
		return errors.New("Missing parameter, recieved parameters: hub.callback:" + cb + " hub.mode: " + mode + " hub.topic: " + topic)
	} else {
		return nil
	}
}

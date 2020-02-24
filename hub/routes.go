package main

import "net/http"

func routes() {
	http.HandleFunc("/", subscriptionHandler)
	http.HandleFunc("/publish", publishingHandler)
}

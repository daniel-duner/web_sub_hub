package main

import (
	"fmt"
	"net/http"
	"os"
)
//Starts up a server, listening to either port defined in enviromental variables or port 8080
func connect() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	fmt.Println(http.ListenAndServe(":"+port, nil))
}

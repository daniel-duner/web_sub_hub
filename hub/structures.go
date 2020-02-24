package main

//A collection defining a subscription
type subscription struct {
	callback string
	mode     string
	topic    string
	secret   string
}

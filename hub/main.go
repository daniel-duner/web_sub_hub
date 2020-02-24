package main

var subscriptions = make(map[string]subscription)

func main() {
	routes()
	connect()
}

/*	SOME LIMITATIONS
	- It is not possible to be subscribed to more than one topic in this solution
	- Subscriptions are stored in memory and will be lost on termination (just a map)
	- The solution does not handle lease_seconds and subscription "time to live"
	- No unit testing, only testing complete transactions through postman & curls
*/

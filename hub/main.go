package main


var subscriptions = make(map[string]subscription)

func main() {
	routes()
	connect()
}

/*	LIMITS
	- It is not possible to be subscribed to more than one topic in this solution
	- Subscriptions are stored in memory and will (just a map)
	- The solution does not handle lease_seconds


*/
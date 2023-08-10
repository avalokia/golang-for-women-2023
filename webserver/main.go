package main

import "webserver/routers"

func main() {
	// Define port
	var PORT = ":9090"
	// Start server
	routers.StartServer().Run(PORT)
}

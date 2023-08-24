package main

import (
	"assignment-project/database"
	"assignment-project/routers"
	"log"
	"os"
)

func main() {
	// Start database
	err := database.StartDB()
	if err != nil {
		log.Fatalln("Error starting DB:", err)
		return
	}
	// Define port
	var PORT = ":" + os.Getenv("WEB_SERVER_PORT")
	// Start server
	routers.StartServer().Run(PORT)
}

package main

import (
	"gauntlet-be/m/v2/database"
	"gauntlet-be/m/v2/internal/handler"

	"log"
)

func main() {

	server := handler.GetRouter()

	// Initiate Database
	database.InitDb()

	err := server.Run("localhost:8080")
	if err != nil {
		log.Print("Error starting server")
		log.Println("Error :", err)
	}
}

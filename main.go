package main

import (
	"log"
	"vijaymehrotra/go-deploy/database"
	"vijaymehrotra/go-deploy/routes"
)

func main() {
	database.InitilizeDB()
	log.Println("Starting server on port 3000...")
	routes.SetupRoutes()
}
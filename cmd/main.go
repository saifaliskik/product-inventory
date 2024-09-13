package main

import (
	"go-product-api/config"
	"go-product-api/routes"
	"log"
)

func main() {
	// Initialize the database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Could not connect to database")
	}
	defer db.Close()

	// Initialize Gin router
	r := routes.SetupRouter(db)

	// Start the server
	r.Run(":8080") // Default port is 8080
}

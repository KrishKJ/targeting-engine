package main

import (
	"log"

	"github.com/KrishKJ/targeting-engine/db"
	"github.com/KrishKJ/targeting-engine/delivery"
	"github.com/gin-gonic/gin"
)

// main function initializes the application
// It connects to the database, sets up the router, and starts the server
func main() {
	// Connect to Postgres and Redis
	db.ConnectPostgres()
	db.ConnectRedis()

	// Auto-migrate DB schema
	db.AutoMigrate()
	// Load campaigns and rules into Redis cache
	db.LoadCampaignsToCache()

	// Create Gin router
	router := gin.Default()

	// Group APIs under /api
	api := router.Group("/api")
	delivery.LoadServices(api)

	// Start the server
	log.Println("🚀 Server started at http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/KrishKJ/targeting-engine/db"
	"github.com/KrishKJ/targeting-engine/delivery"
)

func main() {
	// Connect to Postgres and Redis
	db.ConnectPostgres()
	db.ConnectRedis()

	// Auto-migrate DB schema
	db.AutoMigrate()

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

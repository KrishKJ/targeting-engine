package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	"github.com/KrishKJ/targeting-engine/db"
	"github.com/KrishKJ/targeting-engine/config"
	"github.com/KrishKJ/targeting-engine/delivery"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// main function initializes the application
// It connects to the database, sets up the router, and starts the server
func main() {

	// Load environment variables
	config.LoadEnv()

	// Set GOMAXPROCS to number of logical CPUs
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	log.Printf("🧠 GOMAXPROCS set to %d", numCPU)

	// Start pprof on port 6060 in background
	go func() {
		log.Println("📊 pprof running at http://localhost:6060/debug/pprof/")
		log.Println("📈 Try: go tool pprof http://localhost:6060/debug/pprof/profile")
		http.ListenAndServe("localhost:6060", nil)
	}()

	// Connect to Postgres and Redis
	db.ConnectPostgres()
	db.ConnectRedis()

	// Auto-migrate DB schema
	db.AutoMigrate()

	// Load campaigns and rules into Redis cache
	go func() {
		ticker := time.NewTicker(10 * time.Minute) // Adjust time as needed, currently Redis will refresh every 10 minutes to keep the App updated
		for range ticker.C {
			log.Println("♻️ Refreshing Redis campaign cache...")
			db.LoadCampaignsToCache()
		}
	}()

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
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	log.Println("📊 Prometheus metrics available at /metrics")
}

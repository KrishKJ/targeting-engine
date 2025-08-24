package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"runtime/debug"

	"github.com/KrishKJ/targeting-engine/db"
	"github.com/KrishKJ/targeting-engine/delivery"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	debug.SetGCPercent(20) // Optional performance tuning
	fmt.Printf("🧠 GOMAXPROCS set to %d\n", runtime.GOMAXPROCS(0))
}

func main() {

	// Start pprof server for profiling
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("🔥 panic in pprof server:", r)
			}
		}()
		fmt.Println("📊 pprof running at http://localhost:6060/debug/pprof/")
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	// Connect to DB and Redis
	db.ConnectPostgres()
	db.ConnectRedis()
	db.AutoMigrate()
	db.LoadCampaignsToCache()

	// Setup Gin
	r := gin.Default()

	// 🔥 Expose Prometheus metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Your APIs
	api := r.Group("/api/v1")
	{
		api.GET("/delivery", delivery.HandleDelivery)
		api.GET("/refresh-cache", delivery.RefreshCache)
	}

	// ✅ Optional health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Run server
	fmt.Println("🚀 Server started at http://localhost:8080")
	r.Run(":8080")
}

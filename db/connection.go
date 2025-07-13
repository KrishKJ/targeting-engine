package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"encoding/json"

	"github.com/KrishKJ/targeting-engine/config"
	"github.com/KrishKJ/targeting-engine/delivery/models"
	"github.com/go-redis/redis/v8"
)

// Global variables for DB and Redis connections
var (
	DB    *gorm.DB
	Redis *redis.Client
	Ctx   = context.Background()
)

// ConnectPostgres initializes the Postgres database connection
// Make sure Postgres server is running and the database exists
func ConnectPostgres() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_USER", "postgres"),
		config.GetEnv("DB_PASSWORD", "krishna"),
		config.GetEnv("DB_NAME", "postgres"),
		config.GetEnv("DB_PORT", "5432"),
	)

	var err error
	maxAttempts := 10

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Connected to Postgres")
			return
		}
		log.Printf("❌ Attempt %d: Failed to connect to Postgres: %v", attempts, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatal("❌ Could not connect to Postgres after retries:", err)
}


// ConnectRedis initializes the Redis client
// Make sure Redis server is running on localhost:6379
func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.GetEnv("REDIS_HOST", "localhost") + ":" + config.GetEnv("REDIS_PORT", "6379"),
		Password: "", // No password for now
		DB:       0,
	})
	_, err := Redis.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("❌ Failed to connect to Redis:", err)
	}
	log.Println("✅ Connected to Redis")
}

// AutoMigrate runs migrations for the Campaign and TargetingRule models
func AutoMigrate() {
	err := DB.AutoMigrate(&models.Campaign{}, &models.TargetingRule{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	log.Println("✅ DB Auto-migrated")
}

// LoadCampaignsToCache loads active campaigns and their targeting rules into Redis cache
// This should be called after connecting to Redis and Postgres
func LoadCampaignsToCache() {
	var campaigns []models.Campaign
	var rules []models.TargetingRule
	var cacheData []models.CachedCampaign

	DB.Where("status = ?", "ACTIVE").Find(&campaigns)
	DB.Find(&rules)

	for _, camp := range campaigns {
		cc := models.CachedCampaign{
			CID:     camp.Code,
			Img:     camp.ImageURL,
			CTA:     camp.CTA,
			Include: map[string][]string{},
			Exclude: map[string][]string{},
		}

		for _, rule := range rules {
			if rule.CampaignID == camp.ID {
				if rule.Type == "include" {
					cc.Include[rule.Dimension] = append(cc.Include[rule.Dimension], rule.Value)
				} else {
					cc.Exclude[rule.Dimension] = append(cc.Exclude[rule.Dimension], rule.Value)
				}
			}
		}

		cacheData = append(cacheData, cc)
	}

	data, err := json.Marshal(cacheData)
	if err != nil {
		log.Println("❌ Failed to marshal cache:", err)
		return
	}

	err = Redis.Set(Ctx, "campaigns:active", data, 0).Err()
	if err != nil {
		log.Println("❌ Failed to set Redis cache:", err)
	} else {
		log.Println("✅ Redis cache set with active campaigns")
	}
}

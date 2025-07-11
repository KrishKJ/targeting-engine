// package db

// import (
// 	"context"
// 	// "fmt"
// 	"log"
// 	// "time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/KrishKJ/targeting-engine/models"
// )

// var (
// 	DB     *gorm.DB
// 	Redis  *redis.Client
// 	Ctx    = context.Background()
// )

// func ConnectPostgres() {
// 	// dsn := "host=localhost user=admin password=password dbname=targetingdb port=5432 sslmode=disable"
// 	dsn := "host=localhost user=postgres password=postgres dbname=targetingdb port=5432 sslmode=disable"

// 	var err error
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("❌ Failed to connect to Postgres:", err)
// 	}
// 	log.Println("✅ Connected to Postgres")
// }

// func ConnectRedis() {
// 	Redis = redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password
// 		DB:       0,  // default DB
// 	})

// 	_, err := Redis.Ping(Ctx).Result()
// 	if err != nil {
// 		log.Fatal("❌ Failed to connect to Redis:", err)
// 	}
// 	log.Println("✅ Connected to Redis")
// }

// func AutoMigrate() {
// 	err := DB.AutoMigrate(&models.Campaign{}, &models.TargetingRule{})
// 	if err != nil {
// 		log.Fatal("❌ Migration failed:", err)
// 	}
// 	log.Println("✅ DB Auto-migrated")
// }

package db

import (
	"context"
	// "fmt"
	"log"
	// "time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/KrishKJ/targeting-engine/delivery/models"
	"github.com/go-redis/redis/v8"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
	Ctx   = context.Background()
)

func ConnectPostgres() {
	// dsn := "host=localhost user=postgres password=postgres dbname=targetingdb port=5432 sslmode=disable"
	dsn := "host=localhost user=postgres password=krishna dbname=postgres port=5432 sslmode=disable"

	var err error
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Failed to connect to Postgres:", err)
	}
	log.Println("✅ Connected to Postgres")
}

func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := Redis.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("❌ Failed to connect to Redis:", err)
	}
	log.Println("✅ Connected to Redis")
}

func AutoMigrate() {
	err := DB.AutoMigrate(&models.Campaign{}, &models.TargetingRule{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	log.Println("✅ DB Auto-migrated")
}

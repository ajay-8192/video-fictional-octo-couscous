package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var MongoDB *mongo.Client
var RedisClient *redis.Client

func ConnectPostgresDatabase() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=mydb port=5432 sslmode=disable"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Failed to connect to database", err)
		panic("Failed to connect database")
	}

	println("Database connected successfully!")
}

func ConnectMongoDatabase() {
	uri := ""

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic("Failed to connect mongo database")
	}

	if err := client.Ping(ctx, nil); err != nil {
        panic("Failed to connect mongo database")
    }

    MongoDB = client
    log.Println("Connected to MongoDB successfully")
}

func ConnectRedisClient() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		//Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	defer redisClient.Close()
}

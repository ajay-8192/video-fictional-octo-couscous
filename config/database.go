package config

import (
	"chaterminal/models"
	"context"
	"fmt"
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

	DB_HOST 	:= os.Getenv("DB_HOST")
	DB_USER 	:= os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME 	:= os.Getenv("DB_NAME")
	DB_PORT 	:= os.Getenv("DB_PORT")
	SSL_MODE 	:= os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, SSL_MODE)

	// dsn := "host=localhost user=user password=password dbname=mydb port=5432 sslmode=disable"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Failed to connect to database", err)
		panic("Failed to connect database")
	}

	if err := DB.AutoMigrate(
		&models.User{},
	); err != nil {
		println("Failed to migrate database", err.Error())
		panic("Failed to migrate database")
	}

	println("Database connected successfully!")
}

func ConnectMongoDatabase() {
	uri := os.Getenv("MONGO_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
		Addr: os.Getenv("REDIS_URL"),
		DB:   0,
	})
	defer redisClient.Close()
	log.Println("Connected to Redis successfully")
}

package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

// ConnectRedis initializes the Redis client
func ConnectRedis() {
	redisURL := os.Getenv("REDIS_URL")
	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	// Test the connection
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis!")
}

// CloseRedis closes the Redis connection
func CloseRedis() {
	if err := RedisClient.Close(); err != nil {
		log.Printf("Failed to close Redis: %v", err)
	}
}

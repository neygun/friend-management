package util

import "github.com/redis/go-redis/v9"

// NewRedisClient creates a new Redis client
func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

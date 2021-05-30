package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

// Database interface: defines methods used to handle Redis operations
type Database interface {
	Set(ctx context.Context, key string, value string) (string, error)
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) (string, error)
	Expire(ctx context.Context, key string, duration time.Duration) (string, error)
}

// CreateRedisDatabase initializes the redis database service
func CreateRedisDatabase() (Database, error) {
	var redisUrl = os.Getenv("REDIS_URL")
	if redisUrl == "" {
		fmt.Println("[WARNING] REDIS_URL env variable NOT SET. Defaulting to :6379")
		redisUrl = ":6379"
	}

	// initialize Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// attempt to ping client, panic if connection is not ok
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to ping redis after starting: %s", err.Error()))
	}

	// return client
	return &RedisDatabase{client: client}, nil
}

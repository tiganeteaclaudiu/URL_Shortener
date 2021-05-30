package database

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type RedisDatabase struct {
	client *redis.Client
}

// Set attempts to do SET operation on redis cache
func (r *RedisDatabase) Set(ctx context.Context, key string, value string) (string, error) {
	_, err := r.client.Set(ctx, key, value, 0).Result()
	if err != nil {
		return generateError("set", err)
	}
	return key, nil
}

// Get attempts to do GET operation on redis cache
func (r *RedisDatabase) Get(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return generateError("get", err)
	}
	return value, nil

}

// Delete attempts to do DELETE operation on redis cache
func (r *RedisDatabase) Delete(ctx context.Context, key string) (string, error) {
	_, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return generateError("delete", err)
	}
	return key, nil
}

// Expire attempts to do EXPIRE operation on redis cache
func (r *RedisDatabase) Expire(ctx context.Context, key string, duration time.Duration) (string, error) {
	_, err := r.client.Expire(ctx, key, duration).Result()
	if err != nil {
		return generateError("expire", err)
	}
	return key, nil
}

// generateError generates error based on operation and error message
// Possible errors can be an operation error (ex. entry not found for key) or database failure.
func generateError(operation string, err error) (string, error) {
	if err == redis.Nil {
		return "", errors.Wrap(err, fmt.Sprintf("%s operation failed", operation))
	}

	return "", errors.Wrap(err, "Fatal database failure")
}

package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func initializeRedisClient() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	return rdb
}

func redisSetStatus(jobID string, status string) error {
	rediskey := fmt.Sprintf("job:%s", jobID)
	return rdb.Set(ctx, rediskey, status, 0).Err()
}

func redisGetStatus(jobID string) (value string, err error) {

	rediskey := fmt.Sprintf("job:%s", jobID)
	value, err = rdb.Get(ctx, rediskey).Result()
	return value, err
}

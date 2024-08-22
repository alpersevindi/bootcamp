package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Could not connect to Redis:", err)
	}

	return rdb
}

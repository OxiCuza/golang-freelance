package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type redisClient struct {
	rdb *redis.Client
}

type Redis interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
}

func InitRedis() Redis {
	err := godotenv.Load(".env")
	PanicIfError(err)

	var (
		REDIS_ADDR = os.Getenv("REDIS_ADDR")
		REDIS_PASS = os.Getenv("REDIS_PASS")
	)

	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		Password: REDIS_PASS,
		DB:       0,
	})

	return &redisClient{rdb: client}
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}) error {
	bytes, err := json.Marshal(value)
	PanicIfError(err)

	err = r.rdb.Set(ctx, key, bytes, time.Duration(10)*time.Second).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	response, err := r.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Println("key does not exist")
		return "", err
	} else if err != nil {
		fmt.Println(err)
		return "", err
	}

	return response, nil
}

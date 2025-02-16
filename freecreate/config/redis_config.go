package config

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func RedisConfig(ctx context.Context) *redis.Client{
	rDB, dErr := strconv.Atoi(os.Getenv("REDIS_DB"))
	if dErr != nil {
		panic("could not convert redis db to int")
	}

	rPC, pErr := strconv.Atoi(os.Getenv("REDIS_PROTOCOL"))
	if pErr != nil {
		panic("could not convert redis protocol to int")
	}
	

	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB: rDB,
		Protocol: rPC,
	})

	sErr := client.Set(ctx, "test", "test", 0).Err()
	if sErr != nil {
		panic("error setting Redis test string")
	}

	fmt.Println("successfully connected to Redis!")

	return client
}
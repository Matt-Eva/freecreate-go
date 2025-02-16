package main

import (
	"context"
	"fmt"

	"freecreate/config"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	eErr := godotenv.Load()
	ctx := context.Background()
	if eErr != nil {
		fmt.Println(eErr)
	}

	pg := config.PGConfig()

	mongo := config.MongoConfig(ctx)

	redis := config.RedisConfig(ctx)

	router(pg, mongo, redis)
}

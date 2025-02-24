package main

import (
	"context"
	"fmt"
	"log"

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

	err := MigratePG(pg)
	if err != nil {
		log.Fatal(err)
	}

	router(pg, mongo, redis)
}

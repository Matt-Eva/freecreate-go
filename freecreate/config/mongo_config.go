package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func MongoConfig(ctx context.Context) (*mongo.Client) {
	uri := os.Getenv("MONGO_URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatalf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("successfully connected to Mongo!")

	return client
}
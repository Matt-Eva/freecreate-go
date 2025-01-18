package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
	eErr := godotenv.Load()
	ctx := context.Background()
	if eErr != nil {
		fmt.Println(eErr)
	}
	r := gin.Default()

  	uri := os.Getenv("MONGO_URI")
	  fmt.Println(uri)

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

	fmt.Println("successfull connected to Mongo!")

	PGUser := os.Getenv("PG_USER")
	PGPassword := os.Getenv("PG_PASSWORD")
	PGName := os.Getenv("PG_NAME")
	PGHost := os.Getenv("PG_HOST")
	PGPort := os.Getenv("PG_PORT")
	pgSSL := os.Getenv("PG_SSL")

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		PGUser, PGPassword, PGHost, PGPort, PGName, pgSSL)

		db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to postgres", err)
	}

	fmt.Println("successfully connected to Postgres")
	

	r.GET("/ping", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
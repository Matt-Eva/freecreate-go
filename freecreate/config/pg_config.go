package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func PGConfig() (*sql.DB) {
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

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to postgres", err)
	}

	fmt.Println("successfully connected to Postgres!")

	return db
}
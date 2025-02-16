package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// You can configure the allowed origins here
		frontendOrigin := os.Getenv("FRONTEND_URL")
		if origin == frontendOrigin {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		// Handle preflight requests (OPTIONS)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}


func router(pg *sql.DB, mongo *mongo.Client, redis *redis.Client) {
	r := gin.Default()
	r.Use(CORS())

	// cookieSecret := os.Getenv("COOKIE_SECRET")
	// store := cookie.NewStore([]byte(cookieSecret))

	// r.Use(sessions.Sessions("cookiesesh",store))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
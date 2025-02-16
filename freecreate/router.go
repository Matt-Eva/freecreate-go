package main

import (
	"database/sql"
	"freecreate/config"
	"os"

	"github.com/gin-contrib/sessions"
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
			c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type")
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
	r.Use(sessions.Sessions("mysession", config.SessionConfig()))

	r.GET("/me", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil || user == false {
			c.JSON(401, gin.H{"user": false})
		} else if user == true{
			c.JSON(200, gin.H{"user": true})
		}
	})

	r.POST("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user", true)
		session.Save()

		c.JSON(200, gin.H{"user": true})
	})

	r.DELETE("/logout", func(c *gin.Context){
		session := sessions.Default(c)
		session.Delete("user")
		session.Save()
		
		c.JSON(201, gin.H{"user": false})
	})

	r.Run()
}

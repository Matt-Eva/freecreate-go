package config

import (
	"os"

	"github.com/gin-contrib/sessions/redis"
)

func SessionConfig() redis.Store{
	seshSecret := os.Getenv("SESSION_SECRET")
	store, _ := redis.NewStore(10, "tcp", os.Getenv("REDIS_ADDR"), "", []byte(seshSecret))
	return store
}
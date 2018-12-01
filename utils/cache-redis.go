package utils

import (
	"log"

	"github.com/go-redis/redis"
)

// Cache Configuration Struct
type redisConfig struct {
	Host     string
	Port     string
	Password string
	Name     string
}

// Cache Configuration Variable
var redisCfg redisConfig

// Cache Connection Variable
var Redis *redis.Client

// Cache Connect Function
func redisConnect() *redis.Client {
	// Set Default Redis Cache Name if Empty
	if len(redisCfg.Name) == 0 {
		redisCfg.Name = 0
	}

	// Get Cache Connection
	cache := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.Name,
	})

	// Test Database Connection
	pong, err := cache.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	// Return Current Connection
	return cache
}

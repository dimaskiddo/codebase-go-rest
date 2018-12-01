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
	Name     int
}

// Cache Configuration Variable
var redisCfg redisConfig

// Cache Connection Variable
var Redis *redis.Client

// Cache Connect Function
func redisConnect() *redis.Client {
	// Get Cache Connection
	cache := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.Name,
	})

	// Test Cache Connection
	_, err := cache.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Return Current Connection
	return cache
}

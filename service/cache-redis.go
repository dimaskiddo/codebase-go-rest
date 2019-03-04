package service

import (
	"github.com/go-redis/redis"
)

// Redis Configuration Struct
type redisConfig struct {
	Host     string
	Port     string
	Password string
	Name     int
}

// Redis Configuration Variable
var redisCfg redisConfig

// Redis Variable
var Redis *redis.Client

// Redis Connect Function
func redisConnect() *redis.Client {
	// Initialize Connection
	conn := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.Name,
	})

	// Test Connection
	_, err := conn.Ping().Result()
	if err != nil {
		Log("fatal", "redis-connect", err.Error())
	}

	// Return Connection
	return conn
}

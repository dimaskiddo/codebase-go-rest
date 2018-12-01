package utils

import (
	"strings"
)

// Cache Initialize Function
func initCache() {
	// Cache Configuration Value
	switch strings.ToLower(Config.GetString("CACHE_DRIVER")) {
	case "redis":
		Config.SetDefault("CACHE_PORT", "6379")

		redisCfg.Host = Config.GetString("CACHE_HOST")
		redisCfg.Port = Config.GetString("CACHE_PORT")
		redisCfg.Password = Config.GetString("CACHE_PASSWORD")
		redisCfg.Name = Config.GetString("CACHE_NAME")

		if len(redisCfg.Host) != 0 && len(redisCfg.Port) != 0 &&
			len(redisCfg.Password) != 0 len(redisCfg.Name) != 0 {

			// Do Redis Cache Connection
			Redis = RedisConnect()
		}
	case "memcache":
		Config.SetDefault("CACHE_PORT", "11211")

		memcacheCfg.Host = Config.GetString("CACHE_HOST")
		memcacheCfg.Port = Config.GetString("CACHE_PORT")

		if len(memcacheCfg.Host) != 0 && len(memcacheCfg.Port) != 0 {

			// Do Memcache Cache Connection
			Memcache = memcacheConnect()
		}
	}
}

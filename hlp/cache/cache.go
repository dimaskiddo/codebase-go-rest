package cache

import (
	"strings"

	core "github.com/AasSuhendar/codebase-go-rest/hlp"
)

// Initialize Function in Cache
func init() {
	// Cache Configuration Value
	switch strings.ToLower(core.Config.GetString("CACHE_DRIVER")) {
	case "redis":
		core.Config.SetDefault("CACHE_PORT", "6379")

		redisCfg.Host = core.Config.GetString("CACHE_HOST")
		redisCfg.Port = core.Config.GetString("CACHE_PORT")
		redisCfg.Password = core.Config.GetString("CACHE_PASSWORD")
		redisCfg.Name = core.Config.GetInt("CACHE_NAME")

		if len(redisCfg.Host) != 0 && len(redisCfg.Port) != 0 {

			// Do Redis Cache Connection
			Redis = redisConnect()
		}
	}
}

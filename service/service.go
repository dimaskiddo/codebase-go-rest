package service

import (
	"strings"
)

// Initialize Function in Service
func init() {
	// Initialize Logger
	logInit()

	// Initialize Configuration
	configInit()

	// Initialize Cryptography
	cryptInit()

	// Initialize Database
	if len(strings.ToLower(Config.GetString("DB_DRIVER"))) != 0 {
		dbInit()
	}

	// Initialize Cache
	if len(strings.ToLower(Config.GetString("CACHE_DRIVER"))) != 0 {
		cacheInit()
	}

	// Initialize Storage
	if len(strings.ToLower(Config.GetString("STORAGE_DRIVER"))) != 0 {
		storeInit()
	}

	// Initialize Router
	routerInit()
}

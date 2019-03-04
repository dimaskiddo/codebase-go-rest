package service

import (
	"strings"
)

// Initialize Function in Utils
func Initialize() {
	// Initialize Logger
	initLog()

	// Initialize Configuration
	initConfig()

	// Initialize Cryptography
	initCrypt()

	// Initialize Database
	if len(strings.ToLower(Config.GetString("DB_DRIVER"))) != 0 {
		initDB()
	}

	// Initialize Cache
	if len(strings.ToLower(Config.GetString("CACHE_DRIVER"))) != 0 {
		initCache()
	}

	// Initialize Storage
	if len(strings.ToLower(Config.GetString("STORAGE_DRIVER"))) != 0 {
		initStore()
	}

	// Initialize Router
	initRouter()
}

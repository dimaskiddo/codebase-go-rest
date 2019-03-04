package service

import (
	"strings"
)

// Initialize Function in Utils
func Initialize() {
	// Initialize Logger
	initLog()
	Log("info", "service-initialize", "initialize log")

	// Initialize Configuration
	Log("info", "service-initialize", "initialize config")
	initConfig()

	// Initialize Cryptography
	Log("info", "service-initialize", "initialize crypto")
	initCrypt()

	// Initialize Database
	if len(strings.ToLower(Config.GetString("DB_DRIVER"))) != 0 {
		Log("info", "service-initialize", "initialize database")
		initDB()
	}

	// Initialize Cache
	if len(strings.ToLower(Config.GetString("CACHE_DRIVER"))) != 0 {
		Log("info", "service-initialize", "initialize cache")
		initCache()
	}

	// Initialize Storage
	if len(strings.ToLower(Config.GetString("STORAGE_DRIVER"))) != 0 {
		Log("info", "service-initialize", "initialize storage")
		initStore()
	}

	// Initialize Router
	Log("info", "service-initialize", "initialize router")
	initRouter()
}

package utils

import "strings"

func Bootstrap() {
	// Initialize Configuration
	initConfig()

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
		initStorageS3()
	}

	// Initialize Router
	initRouter()
}

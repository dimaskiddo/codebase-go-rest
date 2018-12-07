package utils

import (
	"log"
	"strings"
)

// Function to Initialize Utils
func Initialize() {
	// Initialize Configuration
	log.Println("Initialize - Config")
	initConfig()

	// Initialize Database
	if len(strings.ToLower(Config.GetString("DB_DRIVER"))) != 0 {
		log.Println("Initialize - Database")
		initDB()
	}

	// Initialize Cache
	if len(strings.ToLower(Config.GetString("CACHE_DRIVER"))) != 0 {
		log.Println("Initialize - Cache")
		initCache()
	}

	// Initialize Storage
	if len(strings.ToLower(Config.GetString("STORAGE_DRIVER"))) != 0 {
		log.Println("Initialize - Storage")
		initStore()
	}

	// Initialize Router
	log.Println("Initialize - Router")
	initRouter()
}

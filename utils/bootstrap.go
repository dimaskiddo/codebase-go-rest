package utils

func Bootstrap() {
	// Initialize Configuration
	initConfig()

	// Initialize Database
	if len(Config.GetString("DB_DRIVER")) != 0 {
		initDB()
	}

	// Initialize Cache
	if len(Config.GetString("CACHE_DRIVER")) != 0 {
		initCache()
	}

	// Initialize Router
	initRouter()
}

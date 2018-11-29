package helpers

import (
	"log"
	"os"
	"strings"

	"github.com/dimaskiddo/frame-go/drivers"

	"github.com/spf13/viper"
)

// Router CORS Configuration Struct
type RouterCORSConfiguration struct {
	Headers []string
	Origins []string
	Methods []string
}

// Router CORS Configuration Variable
var RouterCORS RouterCORSConfiguration

func ConfigInitialize() {
	// Set Configuration Path Value
	configPath := os.Getenv("CONFIG_PATH")
	if len(configPath) == 0 {
		configPath = "./config"
	}

	// Set Configuration File Value
	configFile := os.Getenv("CONFIG_FILE")
	if len(configFile) == 0 {
		configFile = "config"
	}

	// Set Configuration Type Value
	configType := os.Getenv("CONFIG_TYPE")
	if len(configType) == 0 {
		configType = "yaml"
	}

	// Set Configuration Prefix Value
	configPrefix := os.Getenv("CONFIG_PREFIX")
	if len(configPrefix) == 0 {
		configPrefix = "CONFIG"
	}

	// Initialize Configuratior
	cfg := viper.New()

	// Set Configuratior Configuration
	cfg.SetConfigName(configFile)
	cfg.SetConfigType(configType)
	cfg.AddConfigPath(configPath)

	// Set Configurator Environment
	cfg.SetEnvPrefix(configPrefix)
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set Configurator to Auto Bind Configuration Variables to
	// Environment Variables
	cfg.AutomaticEnv()

	// Set Configurator to Load Configuration File
	ConfigLoadFile(cfg)

	// Set Configurator to Set Default Value and
	// Parse Configuration Variables
	ConfigLoadValues(cfg)
}

func ConfigLoadFile(cfg *viper.Viper) {
	// Load Configuration File
	err := cfg.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}

func ConfigLoadValues(cfg *viper.Viper) {
	if cfg != nil {
		// Service IP Value
		cfg.SetDefault("SERVICE_IP", "0.0.0.0")
		ServerConfig.IP = cfg.GetString("SERVICE_IP")

		// Service Port Value
		cfg.SetDefault("SERVICE_PORT", "3000")
		ServerConfig.Port = cfg.GetString("SERVICE_PORT")

		// CORS Allowed Header Value
		cfg.SetDefault("CORS_ALLOWED_HEADER", "X-Requested-With")
		RouterCORS.Headers = []string{cfg.GetString("CORS_ALLOWED_HEADER")}

		// CORS Allowed Origin Value
		cfg.SetDefault("CORS_ALLOWED_ORIGIN", "*")
		RouterCORS.Origins = []string{cfg.GetString("CORS_ALLOWED_ORIGIN")}

		// CORS Allowed Method Value
		cfg.SetDefault("CORS_ALLOWED_METHOD", "'HEAD', 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', 'OPTIONS'")
		RouterCORS.Methods = []string{cfg.GetString("CORS_ALLOWED_METHOD")}

		// JWT Signing Key Value
		cfg.SetDefault("JWT_SIGNING_KEY", "secrets")
		JWTSigningKey = cfg.GetString("JWT_SIGNING_KEY")

		// Database Configuration Value
		// If Database Driver is Set
		switch strings.ToLower(cfg.GetString("DB_DRIVER")) {
		case "mysql":
			cfg.SetDefault("DB_PORT", "3306")
			drivers.MySQLConfig.Host = cfg.GetString("DB_HOST")
			drivers.MySQLConfig.Port = cfg.GetString("DB_PORT")
			drivers.MySQLConfig.User = cfg.GetString("DB_USER")
			drivers.MySQLConfig.Password = cfg.GetString("DB_PASSWORD")
			drivers.MySQLConfig.Name = cfg.GetString("DB_NAME")
		case "mongo":
			cfg.SetDefault("DB_PORT", "27017")
			drivers.MongoConfig.Host = cfg.GetString("DB_HOST")
			drivers.MongoConfig.Port = cfg.GetString("DB_PORT")
			drivers.MongoConfig.User = cfg.GetString("DB_USER")
			drivers.MongoConfig.Password = cfg.GetString("DB_PASSWORD")
			drivers.MongoConfig.Name = cfg.GetString("DB_NAME")
		}
	} else {
		// Log Fatal If Configuratior Not Found
		log.Fatal("Error, Configurator Not Found!")
	}
}

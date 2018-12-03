package utils

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Configurator Variable
var Config *viper.Viper

// Configuration Initialize Function
func initConfig() {
	// Set Configuration Path Value
	configPath := os.Getenv("CONFIG_PATH")
	if len(configPath) == 0 {
		configPath = "./config"
	}

	// Set Configuration File Value
	configFile := os.Getenv("CONFIG_FILE")
	if len(configFile) == 0 {
		configFile = "dev"
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
	Config = viper.New()

	// Set Configuratior Configuration
	Config.SetConfigName(configFile)
	Config.SetConfigType(configType)
	Config.AddConfigPath(configPath)

	// Set Configurator Environment
	Config.SetEnvPrefix(configPrefix)
	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set Configurator to Auto Bind Configuration Variables to
	// Environment Variables
	Config.AutomaticEnv()

	// Set Configurator to Load Configuration File
	configLoadFile()

	// Set Configurator to Set Default Value and
	// Parse Configuration Variables
	configLoadValues()
}

func configLoadFile() {
	// Load Configuration File
	err := Config.ReadInConfig()
	if err != nil {
		log.Println(strings.ToLower(err.Error()))
	}
}

func configLoadValues() {
	// Server IP Value
	Config.SetDefault("SERVER_IP", "0.0.0.0")
	serverCfg.IP = Config.GetString("SERVER_IP")

	// Server Port Value
	Config.SetDefault("SERVER_PORT", "3000")
	serverCfg.Port = Config.GetString("SERVER_PORT")

	// CORS Allowed Header Value
	Config.SetDefault("CORS_ALLOWED_HEADER", "X-Requested-With")
	routerCORSCfg.Headers = []string{Config.GetString("CORS_ALLOWED_HEADER")}

	// CORS Allowed Origin Value
	Config.SetDefault("CORS_ALLOWED_ORIGIN", "*")
	routerCORSCfg.Origins = []string{Config.GetString("CORS_ALLOWED_ORIGIN")}

	// CORS Allowed Method Value
	Config.SetDefault("CORS_ALLOWED_METHOD", "'HEAD', 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', 'OPTIONS'")
	routerCORSCfg.Methods = []string{Config.GetString("CORS_ALLOWED_METHOD")}

	// JWT Signing Key Value
	Config.SetDefault("JWT_SIGNING_KEY", "secrets")
	jwtSigningKey = Config.GetString("JWT_SIGNING_KEY")
}

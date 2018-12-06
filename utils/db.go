package utils

import (
	"strings"
)

// Database Initialize Function
func initDB() {
	// Database Configuration Value
	switch strings.ToLower(Config.GetString("DB_DRIVER")) {
	case "mysql":
		Config.SetDefault("DB_PORT", "3306")

		mysqlCfg.Host = Config.GetString("DB_HOST")
		mysqlCfg.Port = Config.GetString("DB_PORT")
		mysqlCfg.User = Config.GetString("DB_USER")
		mysqlCfg.Password = Config.GetString("DB_PASSWORD")
		mysqlCfg.Name = Config.GetString("DB_NAME")

		if len(mysqlCfg.Host) != 0 && len(mysqlCfg.Port) != 0 &&
			len(mysqlCfg.User) != 0 && len(mysqlCfg.Password) != 0 &&
			len(mysqlCfg.Name) != 0 {

			// Do MySQL Database Connection
			MySQL = mysqlConnect()
		}
	case "mongo":
		Config.SetDefault("DB_PORT", "27017")

		mongoCfg.Host = Config.GetString("DB_HOST")
		mongoCfg.Port = Config.GetString("DB_PORT")
		mongoCfg.User = Config.GetString("DB_USER")
		mongoCfg.Password = Config.GetString("DB_PASSWORD")
		mongoCfg.Name = Config.GetString("DB_NAME")

		if len(mongoCfg.Host) != 0 && len(mongoCfg.Port) != 0 &&
			len(mongoCfg.User) != 0 && len(mongoCfg.Password) != 0 &&
			len(mongoCfg.Name) != 0 {

			// Do Mongo Database Connection
			MongoSession, Mongo = mongoConnect()
		}
	}
}

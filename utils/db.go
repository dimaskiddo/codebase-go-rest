package utils

import (
	"strings"

	"github.com/dimaskiddo/frame-go/dbs"
)

// Database Initialize Function
func InitDB() {
	// Database Configuration Value
	// If Database Driver is Set
	switch strings.ToLower(Config.GetString("DB_DRIVER")) {
	case "mysql":
		Config.SetDefault("DB_PORT", "3306")
		dbs.MySQLConfig.Host = Config.GetString("DB_HOST")
		dbs.MySQLConfig.Port = Config.GetString("DB_PORT")
		dbs.MySQLConfig.User = Config.GetString("DB_USER")
		dbs.MySQLConfig.Password = Config.GetString("DB_PASSWORD")
		dbs.MySQLConfig.Name = Config.GetString("DB_NAME")

		if len(dbs.MySQLConfig.Host) != 0 && len(dbs.MySQLConfig.Port) != 0 &&
			len(dbs.MySQLConfig.User) != 0 && len(dbs.MySQLConfig.Password) != 0 &&
			len(dbs.MySQLConfig.Name) != 0 {

			// Do MySQL Database Connection
			dbs.MySQL = dbs.MySQLConnect()
		}
	case "mongo":
		Config.SetDefault("DB_PORT", "27017")
		dbs.MongoConfig.Host = Config.GetString("DB_HOST")
		dbs.MongoConfig.Port = Config.GetString("DB_PORT")
		dbs.MongoConfig.User = Config.GetString("DB_USER")
		dbs.MongoConfig.Password = Config.GetString("DB_PASSWORD")
		dbs.MongoConfig.Name = Config.GetString("DB_NAME")

		if len(dbs.MongoConfig.Host) != 0 && len(dbs.MongoConfig.Port) != 0 &&
			len(dbs.MongoConfig.User) != 0 && len(dbs.MongoConfig.Password) != 0 &&
			len(dbs.MongoConfig.Name) != 0 {

			// Do Mongo Database Connection
			dbs.MongoSession, dbs.Mongo = dbs.MongoConnect()
		}
	}
}

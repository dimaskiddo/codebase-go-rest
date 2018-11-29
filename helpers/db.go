package helpers

import (
	"strings"

	"github.com/dimaskiddo/frame-go/drivers"
)

// Database Initialize Function
func DBInitialize() {
	// Database Configuration Value
	// If Database Driver is Set
	switch strings.ToLower(Config.GetString("DB_DRIVER")) {
	case "mysql":
		Config.SetDefault("DB_PORT", "3306")
		drivers.MySQLConfig.Host = Config.GetString("DB_HOST")
		drivers.MySQLConfig.Port = Config.GetString("DB_PORT")
		drivers.MySQLConfig.User = Config.GetString("DB_USER")
		drivers.MySQLConfig.Password = Config.GetString("DB_PASSWORD")
		drivers.MySQLConfig.Name = Config.GetString("DB_NAME")

		if len(drivers.MySQLConfig.Host) != 0 && len(drivers.MySQLConfig.Port) != 0 &&
			len(drivers.MySQLConfig.User) != 0 && len(drivers.MySQLConfig.Password) != 0 &&
			len(drivers.MySQLConfig.Name) != 0 {

			// Do MySQL Database Connection
			drivers.MySQLDB = drivers.MySQLConnect()
		}
	case "mongo":
		Config.SetDefault("DB_PORT", "27017")
		drivers.MongoConfig.Host = Config.GetString("DB_HOST")
		drivers.MongoConfig.Port = Config.GetString("DB_PORT")
		drivers.MongoConfig.User = Config.GetString("DB_USER")
		drivers.MongoConfig.Password = Config.GetString("DB_PASSWORD")
		drivers.MongoConfig.Name = Config.GetString("DB_NAME")

		if len(drivers.MongoConfig.Host) != 0 && len(drivers.MongoConfig.Port) != 0 &&
			len(drivers.MongoConfig.User) != 0 && len(drivers.MongoConfig.Password) != 0 &&
			len(drivers.MongoConfig.Name) != 0 {

			// Do Mongo Database Connection
			drivers.MongoSession, drivers.MongoDB = drivers.MongoConnect()
		}
	}
}

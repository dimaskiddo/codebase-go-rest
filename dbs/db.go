package dbs

import (
	"strings"

	"github.com/dimaskiddo/frame-go/utils"
)

// Database Initialize Function
func DatabaseInitialize() {
	// Database Configuration Value
	// If Database Driver is Set
	switch strings.ToLower(utils.Config.GetString("DB_DRIVER")) {
	case "mysql":
		utils.Config.SetDefault("DB_PORT", "3306")
		MySQLConfig.Host = utils.Config.GetString("DB_HOST")
		MySQLConfig.Port = utils.Config.GetString("DB_PORT")
		MySQLConfig.User = utils.Config.GetString("DB_USER")
		MySQLConfig.Password = utils.Config.GetString("DB_PASSWORD")
		MySQLConfig.Name = utils.Config.GetString("DB_NAME")

		if len(MySQLConfig.Host) != 0 && len(MySQLConfig.Port) != 0 &&
			len(MySQLConfig.User) != 0 && len(MySQLConfig.Password) != 0 &&
			len(MySQLConfig.Name) != 0 {

			// Do MySQL Database Connection
			MySQL = MySQLConnect()
		}
	case "mongo":
		utils.Config.SetDefault("DB_PORT", "27017")
		MongoConfig.Host = utils.Config.GetString("DB_HOST")
		MongoConfig.Port = utils.Config.GetString("DB_PORT")
		MongoConfig.User = utils.Config.GetString("DB_USER")
		MongoConfig.Password = utils.Config.GetString("DB_PASSWORD")
		MongoConfig.Name = utils.Config.GetString("DB_NAME")

		if len(MongoConfig.Host) != 0 && len(MongoConfig.Port) != 0 &&
			len(MongoConfig.User) != 0 && len(MongoConfig.Password) != 0 &&
			len(MongoConfig.Name) != 0 {

			// Do Mongo Database Connection
			MongoSesion, Mongo = MongoConnect()
		}
	}
}

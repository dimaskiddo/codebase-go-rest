package db

import (
	"strings"

	"github.com/dimaskiddo/codebase-go-rest/hlp"
)

// Initialize Function in DB
func init() {
	// Database Configuration Value
	switch strings.ToLower(hlp.Config.GetString("DB_DRIVER")) {
	case "mysql":
		hlp.Config.SetDefault("DB_PORT", "3306")

		mysqlCfg.Host = hlp.Config.GetString("DB_HOST")
		mysqlCfg.Port = hlp.Config.GetString("DB_PORT")
		mysqlCfg.User = hlp.Config.GetString("DB_USER")
		mysqlCfg.Password = hlp.Config.GetString("DB_PASSWORD")
		mysqlCfg.Name = hlp.Config.GetString("DB_NAME")

		if len(mysqlCfg.Host) != 0 && len(mysqlCfg.Port) != 0 &&
			len(mysqlCfg.User) != 0 && len(mysqlCfg.Password) != 0 &&
			len(mysqlCfg.Name) != 0 {

			// Do MySQL Database Connection
			MySQL = mysqlConnect()
		}
	case "mongo":
		hlp.Config.SetDefault("DB_PORT", "27017")

		mongoCfg.Host = hlp.Config.GetString("DB_HOST")
		mongoCfg.Port = hlp.Config.GetString("DB_PORT")
		mongoCfg.User = hlp.Config.GetString("DB_USER")
		mongoCfg.Password = hlp.Config.GetString("DB_PASSWORD")
		mongoCfg.Name = hlp.Config.GetString("DB_NAME")

		if len(mongoCfg.Host) != 0 && len(mongoCfg.Port) != 0 &&
			len(mongoCfg.User) != 0 && len(mongoCfg.Password) != 0 &&
			len(mongoCfg.Name) != 0 {

			// Do Mongo Database Connection
			MongoSession, Mongo = mongoConnect()
		}
	}
}

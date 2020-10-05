package db

import (
	"strings"

	"github.com/dimaskiddo/codebase-go-rest/pkg/server"
)

// Initialize Function in DB
func init() {
	// Database Configuration Value
	switch strings.ToLower(server.Config.GetString("DB_DRIVER")) {
	case "mysql":
		server.Config.SetDefault("DB_PORT", "3306")

		mysqlCfg.Host = server.Config.GetString("DB_HOST")
		mysqlCfg.Port = server.Config.GetString("DB_PORT")
		mysqlCfg.User = server.Config.GetString("DB_USER")
		mysqlCfg.Password = server.Config.GetString("DB_PASSWORD")
		mysqlCfg.Name = server.Config.GetString("DB_NAME")

		if len(mysqlCfg.Host) != 0 && len(mysqlCfg.Port) != 0 &&
			len(mysqlCfg.User) != 0 && len(mysqlCfg.Password) != 0 &&
			len(mysqlCfg.Name) != 0 {

			// Do MySQL Database Connection
			MySQL = mysqlConnect()
		}
	case "mongo":
		server.Config.SetDefault("DB_PORT", "27017")

		mongoCfg.Host = server.Config.GetString("DB_HOST")
		mongoCfg.Port = server.Config.GetString("DB_PORT")
		mongoCfg.User = server.Config.GetString("DB_USER")
		mongoCfg.Password = server.Config.GetString("DB_PASSWORD")
		mongoCfg.Name = server.Config.GetString("DB_NAME")

		if len(mongoCfg.Host) != 0 && len(mongoCfg.Port) != 0 &&
			len(mongoCfg.User) != 0 && len(mongoCfg.Password) != 0 &&
			len(mongoCfg.Name) != 0 {

			// Do Mongo Database Connection
			MongoSession, Mongo = mongoConnect()
		}
	}
}

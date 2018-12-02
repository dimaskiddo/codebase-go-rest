package utils

import (
	"database/sql"
	"log"
	"strings"

	// Set Alias for MySQL Driver Package
	// To _ Since This Package Only Used in
	// For MySQL Driver
	_ "github.com/go-sql-driver/mysql"
)

// Database Configuration Struct
type mysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Database Configuration Variable
var mysqlCfg mysqlConfig

// Database Connection Variable
var MySQL *sql.DB

// Database Connect Function
func mysqlConnect() *sql.DB {
	// Get Database Connection
	db, err := sql.Open("mysql", mysqlCfg.User+":"+mysqlCfg.Password+"@tcp("+mysqlCfg.Host+":"+mysqlCfg.Port+")/"+mysqlCfg.Name)
	if err != nil {
		log.Fatal(strings.ToLower(err.Error()))
	}

	// Test Database Connection
	err = db.Ping()
	if err != nil {
		log.Fatal(strings.ToLower(err.Error()))
	}

	// Return Current Connection
	return db
}

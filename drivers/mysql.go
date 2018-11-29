package drivers

import (
	"database/sql"
	"log"

	// Set Alias for MySQL Driver Package
	// To _ Since This Package Only Used in
	// For MySQL Driver
	_ "github.com/go-sql-driver/mysql"
)

// Database Connection Variable
var MySQLDB *sql.DB

// Database Configuration Struct
type MySQLConfiguration struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Database Configuration Variable
var MySQLConfig MySQLConfiguration

// Database Connect Function
func MySQLConnect() *sql.DB {
	// Get Database Connection
	db, err := sql.Open("mysql", MySQLConfig.User+":"+MySQLConfig.Password+"@tcp("+MySQLConfig.Host+":"+MySQLConfig.Port+")/"+MySQLConfig.Name)
	if err != nil {
		log.Fatal(err)
	}

	// Test Database Connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Return Current Connection
	return db
}

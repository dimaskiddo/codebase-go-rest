package utils

import (
	"log"

	// Set Alias for Mongo Driver Package
	// To mgo Since This Package Has Different Name
	// With It's Repository
	mgo "gopkg.in/mgo.v2"
)

// Database Configuration Struct
type mongoConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Database Configuration Variable
var mongoCfg mongoConfig

// Database Connection Variable
var MongoSession *mgo.Session

// Database Connection Variable
var Mongo *mgo.Database

// Database Connect Function
func mongoConnect() (*mgo.Session, *mgo.Database) {
	// Get Session Connection
	sess, err := mgo.Dial(mongoCfg.User + ":" + mongoCfg.Password + "@" + mongoCfg.Host + ":" + mongoCfg.Port + "/" + mongoCfg.Name)
	if err != nil {
		log.Fatal(err)
	}

	// Test Session Connection
	err = sess.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Set Mongo Session to Monotonic
	sess.SetMode(mgo.Monotonic, true)

	// Return Current Session & Database
	return sess, sess.DB(mongoCfg.Name)
}

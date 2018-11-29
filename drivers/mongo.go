package drivers

import (
	"log"

	// Set Alias for Mongo Driver Package
	// To mgo Since This Package Has Different Name
	// With It's Repository
	mgo "gopkg.in/mgo.v2"
)

// Database Connection Variable
var MongoSession *mgo.Session

// Database Connection Variable
var MongoDB *mgo.Database

// Database Configuration Struct
type MongoConfiguration struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Database Configuration Variable
var MongoConfig MongoConfiguration

// Database Connect Function
func MongoConnect() (*mgo.Session, *mgo.Database) {
	// Get Session Connection
	sess, err := mgo.Dial(MongoConfig.User + ":" + MongoConfig.Password + "@" + MongoConfig.Host + ":" + MongoConfig.Port + "/" + MongoConfig.Name)
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
	return sess, sess.DB(MongoConfig.Name)
}

package database

import (
	"github.com/globalsign/mgo"
	"os"
	"log"
)

var db *mgo.Database

// ConnectDatabase establishes connection to database
func ConnectDatabase(connectionURL, databaseName *string) error {

	if connectionURL == nil {
		s := os.Getenv("MONGODB_URL")
		connectionURL = &s
	}
	if databaseName == nil {
		s := os.Getenv("MONGODB_DB")
		databaseName = &s
	}

	session, err := mgo.Dial(*connectionURL)
	if err != nil {
		return err
	}
	db = session.DB(*databaseName)

	// TODO do this properly with right data
	/*
	db.C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
	db.C("users").EnsureIndex(mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	})
	*/


	return err

}

// HealthCheck returns true, if db.Session.Ping is going well
func HealthCheck() bool {
	if db == nil {
		log.Print("db is nil")
		return false
	}

	if db.Session == nil {
		log.Print("db.Session is nil")
		return false
	}

	err := db.Session.Ping()
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}

// GetDatabase returns database connection object
func GetDatabase() *mgo.Database {
	return db
}

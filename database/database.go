package database

import (
	"github.com/globalsign/mgo"
	"os"
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

// GetDatabase returns database connection object
func GetDatabase() *mgo.Database {
	return db
}

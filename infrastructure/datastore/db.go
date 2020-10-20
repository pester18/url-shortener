package datastore

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

func NewDB(dialStr, dbName string) *mgo.Database {
	session, err := mgo.Dial(dialStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db := session.DB(dbName)
	return db
}

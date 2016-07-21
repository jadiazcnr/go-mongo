package mongodb

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// MongoDB implements IDatabase
type MongoDB struct {
	currentDb *mgo.Database
}

//MongoDBSession store mongoDBSession
type MongoDBSession struct {
	*mgo.Session
}

// New set db values
func New() *MongoDB {
	db := &MongoDB{}
	return db
}

//NewSession set mongoDB connection
func (db *MongoDB) NewSession() *MongoDBSession {
	Host := []string{
		"ds015915.mlab.com:15915",
	}
	const (
		Username = "jadiaz"
		Password = "123456"
		Database = "royale"
	)

	Session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		Username: Username,
		Password: Password,
		Database: Database,
	})
	if err != nil {
		panic(err)
	}

	db.currentDb = Session.DB(Database)

	fmt.Printf("Connected to replica set %v!\n", Session.LiveServers())

	return &MongoDBSession{Session}
}

//Insert new document to collection
func (db *MongoDB) Insert(name string, obj interface{}) error {
	return db.currentDb.C(name).Insert(obj)
}

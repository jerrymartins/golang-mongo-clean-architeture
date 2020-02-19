package config

import (
	"gopkg.in/mgo.v2"
	"log"
)

type DAO struct {
	MongoUri string
	Database string
}

var (
	db *mgo.Database
)

func (m *DAO) GetClient() *mgo.Database {
	session, err := mgo.Dial(m.MongoUri)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
	return db
}

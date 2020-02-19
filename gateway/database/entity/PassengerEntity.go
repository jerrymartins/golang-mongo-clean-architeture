package entity

import "gopkg.in/mgo.v2/bson"

type PassengerEntity struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Active      bool          `bson:"active" json:"active"`
}

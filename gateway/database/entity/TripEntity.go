package entity

import (
	"gopkg.in/mgo.v2/bson"
)

type TripEntity struct {
	ID             bson.ObjectId      `bson:"_id" json:"id"`
	Author         string             `bson:"author" json:"author"`
	Model          string             `bson:"model" json:"model"`
	LicensePlate   string             `bson:"licensePlate" json:"licensePlate"`
	AvailablePlace string             `bson:"availablePlace" json:"availablePlace"`
	Type           string             `bson:"type" json:"type"`
	WhatsApp       string             `bson:"whatsapp" json:"whatsapp"`
	Leaving        string             `bson:"leaving" json:"leaving"`
	Destiny        string             `bson:"destiny" json:"destiny"`
	Status         bool               `bson:"status" json:"status"`
	ExitTime       int64              `bson:"exitTime" json:"exitTime"`
	Passengers     []*PassengerEntity `bson:"passengers" json:"passengers"`
}

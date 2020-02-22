package usecase

import (
	. "go-rest-mongo-clean-architeture/gateway/database/entity"
	"gopkg.in/mgo.v2/bson"
)

type TripUseCase struct {
	MongoUri string
	Database string
}

const (
	COLLECTION_TRIP = "trips"
)

func init() {
	config.Read()
	dao.Database = config.Database
	dao.MongoUri = config.MongoUri
	db = dao.GetClient()
}

func (m *TripUseCase) GetAll() ([]TripEntity, error) {
	var trips []TripEntity
	err := db.C(COLLECTION_TRIP).Find(bson.M{}).All(&trips)
	return trips, err
}

func (m *TripUseCase) GetByID(id string) (TripEntity, error) {
	var trip TripEntity
	err := db.C(COLLECTION_TRIP).FindId(bson.ObjectIdHex(id)).One(&trip)
	return trip, err
}

func (m *TripUseCase) Create(trip TripEntity) error {
	err := db.C(COLLECTION_TRIP).Insert(&trip)
	return err
}

func (m *TripUseCase) Delete(id string) error {
	err := db.C(COLLECTION_TRIP).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *TripUseCase) Update(id string, trip TripEntity) error {
	err := db.C(COLLECTION_TRIP).UpdateId(bson.ObjectIdHex(id), &trip)
	return err
}

package usecase

import (
	//. "go-rest-mongo/config"
	. "go-rest-mongo/gateway/database/entity"
	//mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TripDAO struct {
	MongoUri string
	Database string
}

//var (
//	//db
//	db     *mgo.Database
//	dao    = DAO{}
//	config = Config{}
//)

const (
	COLLECTION_TRIP = "trips"
)

func init() {
	config.Read()
	dao.Database = config.Database
	dao.MongoUri = config.MongoUri
	db = dao.GetClient()
}

func (m *TripDAO) GetAll() ([]TripEntity, error) {
	var trips []TripEntity
	err := db.C(COLLECTION_TRIP).Find(bson.M{}).All(&trips)
	return trips, err
}

func (m *TripDAO) GetByID(id string) (TripEntity, error) {
	var trip TripEntity
	err := db.C(COLLECTION_TRIP).FindId(bson.ObjectIdHex(id)).One(&trip)
	return trip, err
}

func (m *TripDAO) Create(trip TripEntity) error {
	err := db.C(COLLECTION_TRIP).Insert(&trip)
	return err
}

func (m *TripDAO) Delete(id string) error {
	err := db.C(COLLECTION_TRIP).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *TripDAO) Update(id string, trip TripEntity) error {
	err := db.C(COLLECTION_TRIP).UpdateId(bson.ObjectIdHex(id), &trip)
	return err
}

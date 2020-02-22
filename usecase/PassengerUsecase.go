package usecase

import (
	. "go-rest-mongo-clean-architeture/config"
	. "go-rest-mongo-clean-architeture/gateway/database/config"
	. "go-rest-mongo-clean-architeture/gateway/database/entity"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PassengerDAO struct {
	MongoUri string
	Database string
}

var (
	//db
	db     *mgo.Database
	dao    = DAO{}
	config = Config{}
)

const (
	COLLECTION = "passengers"
)

func init() {
	config.Read()
	dao.Database = config.Database
	dao.MongoUri = config.MongoUri
	db = dao.GetClient()
}

func (m *PassengerDAO) GetAll() ([]PassengerEntity, error) {
	var passengers []PassengerEntity
	err := db.C(COLLECTION).Find(bson.M{}).All(&passengers)
	return passengers, err
}

func (m *PassengerDAO) GetByID(id string) (PassengerEntity, error) {
	var passenger PassengerEntity
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&passenger)
	return passenger, err
}

func (m *PassengerDAO) Create(passenger PassengerEntity) error {
	err := db.C(COLLECTION).Insert(&passenger)
	return err
}

func (m *PassengerDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *PassengerDAO) Update(id string, passenger PassengerEntity) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &passenger)
	return err
}

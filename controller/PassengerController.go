package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	. "go-rest-mongo-clean-architeture/controller/config"
	. "go-rest-mongo-clean-architeture/gateway/database/entity"
	. "go-rest-mongo-clean-architeture/usecase"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var dao = PassengerDAO{}

func GetAllPassengers(w http.ResponseWriter, r *http.Request) {
	passengers, err := dao.GetAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, passengers)
}

func GetPassengerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	passenger, err := dao.GetByID(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Passenger ID")
		return
	}
	RespondWithJson(w, http.StatusOK, passenger)
}

func CreatePassenger(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var passenger PassengerEntity
	if err := json.NewDecoder(r.Body).Decode(&passenger); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	passenger.ID = bson.NewObjectId()
	if err := dao.Create(passenger); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, passenger)
}

func UpdatePassenger(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var passenger PassengerEntity
	if err := json.NewDecoder(r.Body).Decode(&passenger); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], passenger); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": passenger.Name + " atualizado com sucesso!"})
}

func DeletePassenger(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

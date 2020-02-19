package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	. "go-rest-mongo/gateway/database/entity"
	. "go-rest-mongo/usecase"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var dao = PassengerDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAllPassengers(w http.ResponseWriter, r *http.Request) {
	passengers, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, passengers)
}

func GetPassengerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	passenger, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Passenger ID")
		return
	}
	respondWithJson(w, http.StatusOK, passenger)
}

func CreatePassenger(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var passenger PassengerEntity
	if err := json.NewDecoder(r.Body).Decode(&passenger); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	passenger.ID = bson.NewObjectId()
	if err := dao.Create(passenger); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, passenger)
}

func UpdatePassenger(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var passenger PassengerEntity
	if err := json.NewDecoder(r.Body).Decode(&passenger); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], passenger); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": passenger.Name + " atualizado com sucesso!"})
}

func DeletePassenger(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

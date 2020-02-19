package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	. "go-rest-mongo-clean-architeture/gateway/database/entity"
	. "go-rest-mongo-clean-architeture/usecase"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

var tripUseCase = TripUseCase{}

func GetAllTrips(w http.ResponseWriter, r *http.Request) {
	trips, err := tripUseCase.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, trips)
}

func GetTripByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	trip, err := tripUseCase.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid trip ID")
		return
	}
	respondWithJson(w, http.StatusOK, trip)
}

func CreateTrip(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var trip TripEntity
	if err := json.NewDecoder(r.Body).Decode(&trip); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	trip.ID = bson.NewObjectId()
	trip.ExitTime = time.Now().Unix()
	if err := tripUseCase.Create(trip); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, trip)
}

func UpdateTrip(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var trip TripEntity
	if err := json.NewDecoder(r.Body).Decode(&trip); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := tripUseCase.Update(params["id"], trip); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": trip.Author + " atualizado com sucesso!"})
}

func DeleteTrip(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := tripUseCase.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

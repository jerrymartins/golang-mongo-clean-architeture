package main

import (
	"fmt"
	"github.com/gorilla/mux"
	oauthController "go-rest-mongo-clean-architeture/controller"
	passengerController "go-rest-mongo-clean-architeture/controller"
	tripController "go-rest-mongo-clean-architeture/controller"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/loginLocal/{googleToken}", oauthController.HandleLoggin)

	r.HandleFunc("/login", oauthController.HandleGoogleLogin)
	r.HandleFunc("/callback", oauthController.HandleGoogleCallback)

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/v1/passengers", passengerController.GetAllPassengers).Methods("GET")
	api.HandleFunc("/v1/passengers/{id}", passengerController.GetPassengerByID).Methods("GET")
	api.HandleFunc("/v1/passengers", passengerController.CreatePassenger).Methods("POST")
	api.HandleFunc("/v1/passengers/{id}", passengerController.UpdatePassenger).Methods("PUT")
	api.HandleFunc("/v1/passengers/{id}", passengerController.DeletePassenger).Methods("DELETE")

	api.HandleFunc("/v1/trips", tripController.GetAllTrips).Methods("GET")
	api.HandleFunc("/v1/trips/{id}", tripController.GetTripByID).Methods("GET")
	api.HandleFunc("/v1/trips", tripController.CreateTrip).Methods("POST")
	api.HandleFunc("/v1/trips/{id}", tripController.UpdateTrip).Methods("PUT")
	api.HandleFunc("/v1/trips/{id}", tripController.DeleteTrip).Methods("DELETE")

	api.Use(oauthController.AuthMiddleware)

	var port = ":5000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"milton/helpers"
	"milton/libs/config"
	"milton/libs/mqtt"
	"milton/models"
)

var Config config.Configuration = config.Read()

var m mqtt.MQTT
var db models.DB

func main() {
	m.Setup(Config.MQTT)
	db.Setup()

	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/query-active-units", helpers.WrapHandler(db, queryActiveUnits)).Methods("GET")

	// units
	api.HandleFunc("/get-all-units", helpers.WrapHandler(db, getAllUnits)).Methods("GET")
	api.HandleFunc("/pair-unit", helpers.WrapHandler(db, pairUnit)).Methods("POST")
	api.HandleFunc("/unpair-unit", helpers.WrapHandler(db, unpairUnit)).Methods("POST")

	// pots
	api.HandleFunc("/add-pot", helpers.WrapHandler(db, addPot)).Methods("POST")
	api.HandleFunc("/get-pots/{UnitID}", helpers.WrapHandler(db, getPots)).Methods("GET")
	api.HandleFunc("/update-pot", helpers.WrapHandler(db, updatePot)).Methods("POST")
	api.HandleFunc("/remove-pot", helpers.WrapHandler(db, removePot)).Methods("POST")

	// watering jobs
	api.HandleFunc("/add-job", helpers.WrapHandler(db, addJob)).Methods("POST")
	api.HandleFunc("/remove-job", helpers.WrapHandler(db, removeJob)).Methods("POST")
	api.HandleFunc("/update-job", helpers.WrapHandler(db, updateJob)).Methods("POST")
	api.HandleFunc("/get-jobs", helpers.WrapHandler(db, getJobs)).Methods("GET")
	api.HandleFunc("/get-job/{JobID}", helpers.WrapHandler(db, getJob)).Methods("GET")

	log.Fatal(http.ListenAndServe(Config.Server.Address, router))
}

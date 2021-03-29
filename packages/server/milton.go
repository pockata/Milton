package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Config Configuration = readConfig()

var m MQTT
var db DB

func main() {
	m.setup()
	db.setup()

	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/query-active-units", queryActiveUnits).Methods("GET")

	// units
	api.HandleFunc("/get-all-units", getAllUnits).Methods("GET")
	api.HandleFunc("/pair-unit", pairUnit).Methods("POST")
	api.HandleFunc("/unpair-unit", unpairUnit).Methods("POST")

	// pots
	api.HandleFunc("/add-pot", addPot).Methods("POST")
	api.HandleFunc("/get-pots/{UnitID}", getPots).Methods("GET")
	api.HandleFunc("/update-pot", updatePot).Methods("POST")
	api.HandleFunc("/remove-pot", removePot).Methods("POST")

	// watering jobs
	api.HandleFunc("/add-job", addJob).Methods("POST")
	api.HandleFunc("/remove-job", removeJob).Methods("POST")
	api.HandleFunc("/update-job", updateJob).Methods("POST")
	api.HandleFunc("/get-jobs", getJobs).Methods("GET")
	api.HandleFunc("/get-job/{JobID}", getJob).Methods("GET")

	log.Fatal(http.ListenAndServe(Config.Server.Address, router))
}

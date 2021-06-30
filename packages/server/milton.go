package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"milton/helpers"
	"milton/libs/config"
	"milton/libs/mqtt"
	"milton/models"
	"milton/routes"
)

func main() {
	var m mqtt.MQTT
	var db models.DB

	Config := config.Read()

	db.Setup()
	m.Setup(Config.MQTT)

	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api/").Subrouter()

	api.Use(helpers.CORSHeaders(api, Config.CORS))

	w := helpers.CreateAPIWrapHandler(db)

	api.HandleFunc("/query-active-units", w(routes.QueryActiveUnits)).Methods("GET")

	// units
	api.HandleFunc("/get-paired-units", w(routes.GetPairedUnits)).Methods("GET")
	api.HandleFunc("/pair-unit", w(routes.PairUnit)).Methods("POST")
	api.HandleFunc("/unpair-unit", w(routes.UnpairUnit)).Methods("POST")

	// pots
	api.HandleFunc("/add-pot", w(routes.AddPot)).Methods("POST")
	api.HandleFunc("/get-pots/{UnitID}", w(routes.GetPots)).Methods("GET")
	api.HandleFunc("/update-pot", w(routes.UpdatePot)).Methods("POST")
	api.HandleFunc("/remove-pot", w(routes.RemovePot)).Methods("POST")

	// watering jobs
	api.HandleFunc("/add-job", w(routes.AddJob)).Methods("POST")
	api.HandleFunc("/remove-job", w(routes.RemoveJob)).Methods("POST")
	api.HandleFunc("/update-job", w(routes.UpdateJob)).Methods("POST")
	api.HandleFunc("/get-jobs", w(routes.GetJobs)).Methods("GET")
	api.HandleFunc("/get-job/{JobID}", w(routes.GetJob)).Methods("GET")

	log.Fatal(http.ListenAndServe(Config.Server.Address, router))
}

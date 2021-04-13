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

	api.HandleFunc("/query-active-units", helpers.WrapHandler(db, routes.QueryActiveUnits)).Methods("GET")

	// units
	api.HandleFunc("/get-all-units", helpers.WrapHandler(db, routes.GetAllUnits)).Methods("GET")
	api.HandleFunc("/pair-unit", helpers.WrapHandler(db, routes.PairUnit)).Methods("POST")
	api.HandleFunc("/unpair-unit", helpers.WrapHandler(db, routes.UnpairUnit)).Methods("POST")

	// pots
	api.HandleFunc("/add-pot", helpers.WrapHandler(db, routes.AddPot)).Methods("POST")
	api.HandleFunc("/get-pots/{UnitID}", helpers.WrapHandler(db, routes.GetPots)).Methods("GET")
	api.HandleFunc("/update-pot", helpers.WrapHandler(db, routes.UpdatePot)).Methods("POST")
	api.HandleFunc("/remove-pot", helpers.WrapHandler(db, routes.RemovePot)).Methods("POST")

	// watering jobs
	api.HandleFunc("/add-job", helpers.WrapHandler(db, routes.AddJob)).Methods("POST")
	api.HandleFunc("/remove-job", helpers.WrapHandler(db, routes.RemoveJob)).Methods("POST")
	api.HandleFunc("/update-job", helpers.WrapHandler(db, routes.UpdateJob)).Methods("POST")
	api.HandleFunc("/get-jobs", helpers.WrapHandler(db, routes.GetJobs)).Methods("GET")
	api.HandleFunc("/get-job/{JobID}", helpers.WrapHandler(db, routes.GetJob)).Methods("GET")

	log.Fatal(http.ListenAndServe(Config.Server.Address, router))
}

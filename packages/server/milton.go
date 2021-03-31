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

type Milton struct {
	Router *mux.Router
	DB     models.DB
	Config config.Configuration
	MQTT   mqtt.MQTT
}

func (m *Milton) Initialize(Config config.Configuration) {
	m.Config = Config

	m.DB.Setup()
	m.MQTT.Setup(m.Config.MQTT)

	m.Router = mux.NewRouter().StrictSlash(true)

	api := m.Router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/query-active-units", helpers.WrapHandler(m.DB, routes.QueryActiveUnits)).Methods("GET")

	// units
	api.HandleFunc("/get-all-units", helpers.WrapHandler(m.DB, routes.GetAllUnits)).Methods("GET")
	api.HandleFunc("/pair-unit", helpers.WrapHandler(m.DB, routes.PairUnit)).Methods("POST")
	api.HandleFunc("/unpair-unit", helpers.WrapHandler(m.DB, routes.UnpairUnit)).Methods("POST")

	// pots
	api.HandleFunc("/add-pot", helpers.WrapHandler(m.DB, routes.AddPot)).Methods("POST")
	api.HandleFunc("/get-pots/{UnitID}", helpers.WrapHandler(m.DB, routes.GetPots)).Methods("GET")
	api.HandleFunc("/update-pot", helpers.WrapHandler(m.DB, routes.UpdatePot)).Methods("POST")
	api.HandleFunc("/remove-pot", helpers.WrapHandler(m.DB, routes.RemovePot)).Methods("POST")

	// watering jobs
	api.HandleFunc("/add-job", helpers.WrapHandler(m.DB, routes.AddJob)).Methods("POST")
	api.HandleFunc("/remove-job", helpers.WrapHandler(m.DB, routes.RemoveJob)).Methods("POST")
	api.HandleFunc("/update-job", helpers.WrapHandler(m.DB, routes.UpdateJob)).Methods("POST")
	api.HandleFunc("/get-jobs", helpers.WrapHandler(m.DB, routes.GetJobs)).Methods("GET")
	api.HandleFunc("/get-job/{JobID}", helpers.WrapHandler(m.DB, routes.GetJob)).Methods("GET")
}

func (m *Milton) Run() {
	log.Fatal(http.ListenAndServe(m.Config.Server.Address, m.Router))
}

func main() {
	var milton Milton

	config := config.Read()

	milton.Initialize(config)
}

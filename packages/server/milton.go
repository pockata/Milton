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

	api.HandleFunc("/query-active-units", w(routes.QueryActiveUnits)).Methods(http.MethodGet)

	// units
	api.HandleFunc("/get-paired-units", w(routes.GetPairedUnits)).Methods(http.MethodGet)
	api.HandleFunc("/pair-unit", w(routes.PairUnit)).Methods(http.MethodPost)
	api.HandleFunc("/unpair-unit", w(routes.UnpairUnit)).Methods(http.MethodPost)

	// pots
	api.HandleFunc("/add-pot", w(routes.AddPot)).Methods(http.MethodPost)
	api.HandleFunc("/get-pots/{UnitID}", w(routes.GetPots)).Methods(http.MethodGet)
	api.HandleFunc("/update-pot", w(routes.UpdatePot)).Methods(http.MethodPost)
	api.HandleFunc("/remove-pot", w(routes.RemovePot)).Methods(http.MethodPost)

	// watering jobs
	api.HandleFunc("/add-job", w(routes.AddJob)).Methods(http.MethodPost)
	api.HandleFunc("/remove-job", w(routes.RemoveJob)).Methods(http.MethodPost)
	api.HandleFunc("/update-job", w(routes.UpdateJob)).Methods(http.MethodPost)
	api.HandleFunc("/get-jobs", w(routes.GetJobs)).Methods(http.MethodGet)
	api.HandleFunc("/get-job/{JobID}", w(routes.GetJob)).Methods(http.MethodGet)

	// TODO: Support PORT env variable
	// httpPort := os.Getenv("PORT")
	// 	if httpPort == "" {
	// 		httpPort = "8080"
	// 	}
	log.Fatal(http.ListenAndServe(Config.Server.Address, router))
}

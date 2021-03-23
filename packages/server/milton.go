package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"milton/routes"
)

var Config Configuration = readConfig()

var m MQTT
var db DB

func main() {
	m.setup()
	db.setup()

	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/query-active-units", routes.QueryActiveUnits)

	log.Fatal(http.ListenAndServe(Config.Server.Address, router))
}

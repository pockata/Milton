package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Config Configuration = readConfig()

var m MQTT

func main() {
	m.setup()

	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/query-active-pumps", queryPumps)

	log.Fatal(http.ListenAndServe(Config.Server.Address, router))
}

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

	db.init()
	// defer db.close()
	db.log("lek", "vyiiiiiiiiiii")

	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/query-active-pumps", queryPumps)

	log.Fatal(http.ListenAndServe(Config.Server.Address, router))
}

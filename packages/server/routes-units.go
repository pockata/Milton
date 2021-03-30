package main

import (
	"errors"
	"log"
	"net/http"

	"milton/helpers"
	"milton/models"
)

func getAllUnits(rw http.ResponseWriter, r *http.Request) {
	var units []models.Unit
	db.Instance.Find(&units)

	helpers.SuccessResponse(rw, units)
}

func pairUnit(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data", err)
		return
	}

	mdns := r.PostForm.Get("MDNS")
	name := r.PostForm.Get("Name")

	if !helpers.CheckParams(mdns, name) {
		helpers.ErrorResponse(rw, errors.New("Invalid request. Missing parameters"))
		return
	}

	entry := &models.Unit{MDNS: mdns, Name: name}
	helpers.CreateEntry(rw, r, *db.Instance, &entry)
}

func unpairUnit(rw http.ResponseWriter, r *http.Request) {
	helpers.DeleteEntry(rw, r, *db.Instance, &models.Job{})
}

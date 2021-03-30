package main

import (
	"errors"
	"log"
	"net/http"

	"milton/models"
)

func getAllUnits(rw http.ResponseWriter, r *http.Request) {
	var units []models.Unit
	db.Instance.Find(&units)

	successResponse(rw, units)
}

func pairUnit(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data", err)
		return
	}

	mdns := r.PostForm.Get("MDNS")
	name := r.PostForm.Get("Name")

	if !checkParams(mdns, name) {
		errorResponse(rw, errors.New("Invalid request. Missing parameters"))
		return
	}

	entry := &models.Unit{MDNS: mdns, Name: name}
	createEntry(rw, r, *db.Instance, &entry)
}

func unpairUnit(rw http.ResponseWriter, r *http.Request) {
	deleteEntry(rw, r, *db.Instance, &models.Job{})
}

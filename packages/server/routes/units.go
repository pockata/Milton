package routes

import (
	"errors"
	"log"
	"net/http"

	"milton/helpers"
	"milton/models"
)

func GetPairedUnits(rw http.ResponseWriter, r *http.Request, db models.DB) {
	var units []models.Unit
	db.Instance.Find(&units)

	helpers.SuccessResponse(rw, r, units)
}

func PairUnit(rw http.ResponseWriter, r *http.Request, db models.DB) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data", err)
		return
	}

	mdns := r.PostForm.Get("MDNS")
	name := r.PostForm.Get("Name")

	if !helpers.CheckParams(mdns, name) {
		helpers.ErrorResponse(rw, r, errors.New("Invalid request. Missing parameters"))
		return
	}

	entry := &models.Unit{MDNS: mdns, Name: name}
	helpers.CreateEntry(rw, r, *db.Instance, &entry)
}

func UnpairUnit(rw http.ResponseWriter, r *http.Request, db models.DB) {
	helpers.DeleteEntry(rw, r, *db.Instance, &models.Job{})
}

package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"milton/helpers"
	"milton/models"
)

func addPot(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data", err)
		return
	}

	unitID := r.PostForm.Get("UnitID")
	name := r.PostForm.Get("Name")

	if !helpers.CheckParams(name, unitID) {
		helpers.ErrorResponse(rw, errors.New("Invalid request. Missing parameters"))
		return
	}

	var unit models.Unit
	find := db.Instance.First(&unit, unitID)

	if find.Error != nil {
		helpers.ErrorResponse(rw, find.Error)
		return
	}

	entry := &models.Pot{UnitID: unit.ID, Name: name}
	helpers.CreateEntry(rw, r, *db.Instance, &entry)
}

func removePot(rw http.ResponseWriter, r *http.Request) {
	helpers.DeleteEntry(rw, r, *db.Instance, &models.Pot{})
}

func getPots(rw http.ResponseWriter, r *http.Request) {
	var pots []models.Pot
	var unit models.Unit

	vars := mux.Vars(r)

	unitID, err := strconv.Atoi(vars["UnitID"])
	if err != nil {
		helpers.ErrorResponse(rw, errors.New("Invalid unit ID"))
		return
	}

	findUnit := db.Instance.First(&unit, unitID)
	if findUnit.Error != nil {
		helpers.ErrorResponse(rw, errors.New("Non-existing unit ID"))
		return
	}

	db.Instance.Model(&unit).Association("Pots").Find(&pots)

	helpers.SuccessResponse(rw, pots)
}

func updatePot(rw http.ResponseWriter, r *http.Request) {
	var pot []models.Pot

	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data", err)
		return
	}

	potID := r.PostForm.Get("PotID")
	name := r.PostForm.Get("Name")
	UnitIDStr := r.PostForm.Get("UnitID")

	if !helpers.CheckParams(potID, name, UnitIDStr) {
		helpers.ErrorResponse(rw, errors.New("Invalid parameters"))
		return
	}

	UnitID, err := strconv.ParseUint(UnitIDStr, 10, 32)

	if err != nil {
		helpers.ErrorResponse(rw, err)
		return
	}

	findPot := db.Instance.Find(&pot, potID)
	if findPot.Error != nil {
		helpers.ErrorResponse(rw, findPot.Error)
		return
	}

	db.Instance.Model(&pot).Updates(models.Pot{
		Name:   name,
		UnitID: uint(UnitID),
	})

	helpers.SuccessResponse(rw, &pot)
}

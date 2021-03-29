package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func addPot(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data", err)
		return
	}

	unitID := r.PostForm.Get("UnitID")
	name := r.PostForm.Get("Name")

	if !checkParams(name, unitID) {
		errorResponse(rw, errors.New("Invalid request. Missing parameters"))
		return
	}

	var unit Unit
	find := db.instance.First(&unit, unitID)

	if find.Error != nil {
		errorResponse(rw, find.Error)
		return
	}

	entry := &Pot{UnitID: unit.ID, Name: name}
	createEntry(rw, r, *db.instance, &entry)
}

func removePot(rw http.ResponseWriter, r *http.Request) {
	deleteEntry(rw, r, *db.instance, &Pot{})
}

func getPots(rw http.ResponseWriter, r *http.Request) {
	var pots []Pot
	var unit Unit

	vars := mux.Vars(r)

	unitID, err := strconv.Atoi(vars["UnitID"])
	if err != nil {
		errorResponse(rw, errors.New("Invalid unit ID"))
		return
	}

	findUnit := db.instance.First(&unit, unitID)
	if findUnit.Error != nil {
		errorResponse(rw, errors.New("Non-existing unit ID"))
		return
	}

	db.instance.Model(&unit).Association("Pots").Find(&pots)

	successResponse(rw, pots)
}

func updatePot(rw http.ResponseWriter, r *http.Request) {
	var pot []Pot

	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data", err)
		return
	}

	potID := r.PostForm.Get("PotID")
	name := r.PostForm.Get("Name")
	UnitIDStr := r.PostForm.Get("UnitID")

	if !checkParams(potID, name, UnitIDStr) {
		errorResponse(rw, errors.New("Invalid parameters"))
		return
	}

	UnitID, err := strconv.ParseUint(UnitIDStr, 10, 32)

	if err != nil {
		errorResponse(rw, err)
		return
	}

	findPot := db.instance.Find(&pot, potID)
	if findPot.Error != nil {
		errorResponse(rw, findPot.Error)
		return
	}

	db.instance.Model(&pot).Updates(Pot{
		Name:   name,
		UnitID: uint(UnitID),
	})

	successResponse(rw, &pot)
}

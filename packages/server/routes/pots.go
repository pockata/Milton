package routes

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/volatiletech/sqlboiler/v4/boil"

	models "milton/generated_models"
	"milton/helpers"
)

type CreatePotResponse struct {
	Pot models.FlowerPot `json:"flowerPot"`
}

func AddPot(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseForm()
	if err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	unitID := r.PostForm.Get("UnitID")
	name := r.PostForm.Get("Name")

	if !helpers.CheckParams(name, unitID) {
		helpers.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	unit, err := models.FindUnit(context.Background(), db, unitID)
	if err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't find unit: %w", err))
		return
	}

	pot := models.FlowerPot{
		UnitID: unit.ID,
		Name:   name,
	}

	if err := pot.Insert(context.Background(), db, boil.Infer()); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error inserting pot: %w", err))
		return
	}

	helpers.SuccessResponse(w, r, CreatePotResponse{
		Pot: pot,
	})
}

type RemovePotResponse struct {
	Success bool `json:"success"`
}

func RemovePot(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if err := r.ParseForm(); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	ID := r.Form.Get("ID")
	if !helpers.CheckParams(ID) {
		helpers.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	pot, err := models.FindFlowerPot(context.Background(), db, ID)
	if err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error finding flower pot: %w", err))
		return
	}

	if _, err := pot.Delete(context.Background(), db); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error deleting flower pot: %w", err))
		return
	}

	helpers.SuccessResponse(w, r, RemovePotResponse{
		Success: true,
	})
}

type GetPotsResponse struct {
	Pots models.FlowerPotSlice `json:"flowerPots"`
}

func GetPots(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	unitID := vars["UnitID"]
	if helpers.CheckParams(unitID) {
		helpers.ErrorResponse(rw, r, fmt.Errorf("invalid unit ID: %v", unitID))
		return
	}

	unit, err := models.FindUnit(context.Background(), db, unitID)
	if err != nil {
		helpers.ErrorResponse(rw, r, fmt.Errorf("couldn't find unit id: %w", err))
		return
	}

	pots, err := unit.UnitFlowerPots().All(context.Background(), db)
	if err != nil {
		helpers.ErrorResponse(rw, r, fmt.Errorf("couldn't get unit pots: %w", err))
		return
	}

	helpers.SuccessResponse(rw, r, GetPotsResponse{
		Pots: pots,
	})
}

type UpdatePotResponse struct {
	Success bool `json:"success"`
}

func UpdatePot(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if err := r.ParseForm(); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	potID := r.PostForm.Get("PotID")
	name := r.PostForm.Get("Name")
	UnitID := r.PostForm.Get("UnitID")

	if !helpers.CheckParams(potID, name, UnitID) {
		helpers.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	pot, err := models.FindFlowerPot(context.Background(), db, potID)
	if err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't find pot: %w", err))
		return
	}

	pot.Name = name
	pot.UnitID = UnitID

	if _, err := pot.Update(context.Background(), db, boil.Infer()); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't update pot: %w", err))
		return
	}

	helpers.SuccessResponse(w, r, UpdatePotResponse{
		Success: true,
	})
}

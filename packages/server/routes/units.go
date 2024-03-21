package routes

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	models "milton/generated_models"
	"milton/helpers"

	"github.com/lucsky/cuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PairedUnitsResponse struct {
	Units models.UnitSlice `json:"units"`
}

func GetPairedUnits(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	var units models.UnitSlice

	units, _ = models.Units().All(context.Background(), db)
	helpers.SuccessResponse(rw, r, PairedUnitsResponse{
		Units: units,
	})
}

type PairUnitResponse struct {
	Unit models.Unit `json:"unit"`
}

func PairUnit(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseForm()
	if err != nil {
		helpers.ErrorResponse(w, r, err)
		log.Println("Error parsing form data", err)
		return
	}

	mdns := r.PostForm.Get("MDNS")
	name := r.PostForm.Get("Name")

	if !helpers.CheckParams(mdns, name) {
		helpers.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	unit := models.Unit{
		ID:   fmt.Sprintf("u-%s", cuid.New()),
		Name: name,
		MDNS: mdns,
	}

	err = unit.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("insertion failed: %w", err))
		return
	}

	helpers.SuccessResponse(w, r, PairUnitResponse{
		Unit: unit,
	})
}

type UnpairUnitResponse struct {
	Success bool `json:"success"`
}

func UnpairUnit(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseForm()
	if err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't parse form data %w", err))
		return
	}

	ID := r.Form.Get("ID")
	if !helpers.CheckParams(ID) {
		helpers.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	unit, err := models.FindUnit(context.Background(), db, ID)

	if err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't find unit to delete: %w", err))
		return
	}

	if _, err := unit.Delete(context.Background(), db); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't delete unit: %w", err))
		return
	}

	helpers.SuccessResponse(w, r, UnpairUnitResponse{
		Success: true,
	})
}

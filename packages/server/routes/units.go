package routes

import (
	"errors"
	"fmt"
	"milton"
	"net/http"
)

type PairedUnitsResponse struct {
	Units milton.UnitSlice `json:"units"`
}

func (c Controller) GetPairedUnits(rw http.ResponseWriter, r *http.Request) {
	units, err := c.app.GetAllUnits()
	if err != nil {
		c.ErrorResponse(rw, r, err)
		return
	}

	c.SuccessResponse(rw, r, PairedUnitsResponse{
		Units: units,
	})
}

type PairUnitResponse struct {
	Unit milton.Unit `json:"unit"`
}

func (c Controller) PairUnit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	mdns := r.PostForm.Get("MDNS")
	name := r.PostForm.Get("Name")

	if !c.ValidParams(name, mdns) {
		c.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	unit, err := c.app.PairUnit(name, mdns)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.SuccessResponse(w, r, PairUnitResponse{
		Unit: unit,
	})
}

type UnpairUnitResponse struct {
	Success bool `json:"success"`
}

func (c Controller) UnpairUnit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	ID := r.Form.Get("ID")
	if !c.ValidParams(ID) {
		c.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	if err := c.app.UnpairUnit(ID); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.SuccessResponse(w, r, UnpairUnitResponse{
		Success: true,
	})
}

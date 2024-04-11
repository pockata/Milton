package http

import (
	"errors"
	"fmt"
	"milton/core/domain"
	"net/http"
)

type PairedUnitsResponse struct {
	Units domain.UnitSlice `json:"units"`
}

func (c HTTPController) GetPairedUnits(rw http.ResponseWriter, r *http.Request) {
	units, err := c.units.GetAll()
	if err != nil {
		c.ErrorResponse(rw, r, err)
		return
	}

	c.SuccessResponse(rw, r, PairedUnitsResponse{
		Units: units,
	})
}

type PairUnitResponse struct {
	Unit domain.Unit `json:"unit"`
}

func (c HTTPController) PairUnit(w http.ResponseWriter, r *http.Request) {
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

	unit, err := c.units.Pair(name, mdns)
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

func (c HTTPController) UnpairUnit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	ID := r.Form.Get("ID")
	if !c.ValidParams(ID) {
		c.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	if err := c.units.Unpair(ID); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.SuccessResponse(w, r, UnpairUnitResponse{
		Success: true,
	})
}

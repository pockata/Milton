package routes

import (
	"errors"
	"fmt"
	"milton"
	"net/http"

	"milton/helpers"
)

type CreatePotResponse struct {
	Pot milton.FlowerPot `json:"flowerPot"`
}

func (c Controller) AddPot(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	unitID := r.PostForm.Get("UnitID")
	name := r.PostForm.Get("Name")

	if !helpers.ValidParams(name, unitID) {
		helpers.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	pot, err := c.app.AddFlowerPot(name, unitID)
	if err != nil {
		helpers.ErrorResponse(w, r, err)
		return
	}

	helpers.SuccessResponse(w, r, CreatePotResponse{
		Pot: pot,
	})
}

type RemovePotResponse struct {
	Success bool `json:"success"`
}

func (c Controller) RemovePot(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	ID := r.Form.Get("ID")
	if !helpers.ValidParams(ID) {
		helpers.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	if err := c.app.RemoveFlowerPot(ID); err != nil {
		helpers.ErrorResponse(w, r, err)
		return
	}

	helpers.SuccessResponse(w, r, RemovePotResponse{
		Success: true,
	})
}

type GetPotsResponse struct {
	FlowerPots milton.FlowerPotSlice `json:"flowerPots"`
}

func (c Controller) GetPots(rw http.ResponseWriter, r *http.Request) {
	unitID := r.PathValue("UnitID")

	if !helpers.ValidParams(unitID) {
		helpers.ErrorResponse(rw, r, fmt.Errorf("invalid unit ID: %v", unitID))
		return
	}

	pots, err := c.app.GetFlowerPots(unitID)
	if err != nil {
		helpers.ErrorResponse(rw, r, err)
		return
	}

	helpers.SuccessResponse(rw, r, GetPotsResponse{
		FlowerPots: pots,
	})
}

type UpdatePotResponse struct {
	Success bool `json:"success"`
}

func (c Controller) RenamePot(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	potID := r.PostForm.Get("PotID")
	name := r.PostForm.Get("Name")

	if !helpers.ValidParams(potID, name) {
		helpers.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	if err := c.app.RenameFlowerPot(potID, name); err != nil {
		helpers.ErrorResponse(w, r, err)
		return
	}

	helpers.SuccessResponse(w, r, UpdatePotResponse{
		Success: true,
	})
}

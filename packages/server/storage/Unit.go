package storage

import (
	"milton"
	models "milton/generated_models"
)

type Unit struct {
	unit *models.Unit
}

func NewUnit(unit *models.Unit) milton.Unit {
	return &Unit{
		unit: unit,
	}
}

func (u *Unit) Name() string {
	return u.unit.Name
}

func (u *Unit) MDNS() string {
	return u.unit.MDNS
}

func (u *Unit) ID() string {
	return u.unit.ID
}

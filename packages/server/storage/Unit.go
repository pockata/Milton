package storage

import (
	"database/sql"
	models "milton/generated_models"
)

type Unit struct {
	unit *models.Unit
	db   *sql.DB
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

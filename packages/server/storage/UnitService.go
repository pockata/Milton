package storage

import (
	"context"
	"database/sql"
	"milton"
	models "milton/generated_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UnitService struct {
	db *sql.DB
}

func (u *UnitService) Pair(mdns string, name string) error {
	entry := &models.Unit{MDNS: mdns, Name: name}
	ctx := context.Background()

	return entry.Insert(ctx, u.db, boil.Infer())
}

func (u *UnitService) Unpair(ID string) error {
	ctx := context.Background()

	unit, err := models.FindUnit(ctx, u.db, ID)

	if err != nil {
		return err
	}

	_, err = unit.Delete(ctx, u.db)

	return err
}

func (u *UnitService) All() ([]milton.Unit, error) {
	ctx := context.Background()

	units, err := models.Units().All(ctx, u.db)

	if err != nil {
		return nil, err
	}

	mUnits := make([]milton.Unit, len(units))

	for i, un := range units {
		mUnits[i] = transformUnit(un, u.db)
	}

	return mUnits, nil
}

func transformUnit(unit *models.Unit, db *sql.DB) milton.Unit {
	return &Unit{
		unit: unit,
		db:   db,
	}
}

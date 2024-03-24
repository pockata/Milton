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

func NewUnitService(db *sql.DB) UnitService {
	return UnitService{
		db: db,
	}
}

func (u *UnitService) Get(ID string) (milton.Unit, error) {
	ctx := context.Background()
	unit, err := models.FindUnit(ctx, u.db, ID)

	if err != nil {
		return nil, err
	}

	return unit, nil
}

func (u *UnitService) GetPots(unit milton.Unit) (milton.FlowerPotSlice, error) {
	pots, err := unit.UnitFlowerPots().All(context.Background(), u.db)
	if err != nil {
		return nil, err
	}

	return pots, nil
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

func (u *UnitService) All() (milton.UnitSlice, error) {
	ctx := context.Background()

	units, err := models.Units().All(ctx, u.db)

	if err != nil {
		return nil, err
	}

	return units, nil
}

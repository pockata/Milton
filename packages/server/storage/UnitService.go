package storage

import (
	"context"
	"database/sql"
	"fmt"
	"milton"
	models "milton/generated_models"

	"github.com/lucsky/cuid"
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

func (u UnitService) Get(ID string) (milton.Unit, error) {
	unit, err := models.FindUnit(context.Background(), u.db, ID)
	if err != nil {
		return nil, err
	}

	return unit, nil
}

func (u UnitService) GetPots(unit milton.Unit) (milton.FlowerPotSlice, error) {
	return unit.UnitFlowerPots().All(context.Background(), u.db)
}

func (u UnitService) Pair(name string, mdns string) (milton.Unit, error) {
	unit := &models.Unit{
		ID:   fmt.Sprintf("u-%s", cuid.New()),
		Name: name,
		MDNS: mdns,
	}

	err := unit.Insert(context.Background(), u.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return unit, nil
}

func (u UnitService) Unpair(ID string) error {
	unit, err := models.FindUnit(context.Background(), u.db, ID)
	if err != nil {
		return err
	}

	_, err = unit.Delete(context.Background(), u.db)

	return err
}

func (u UnitService) All() (milton.UnitSlice, error) {
	units, err := models.Units().All(context.Background(), u.db)

	if err != nil {
		return nil, err
	}

	return units, nil
}

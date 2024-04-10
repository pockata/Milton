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

type UnitRepository struct {
	db *sql.DB
}

func NewUnitRepository(db *sql.DB) UnitRepository {
	return UnitRepository{
		db: db,
	}
}

func (u UnitRepository) Get(ID string) (milton.Unit, error) {
	unit, err := models.FindUnit(context.Background(), u.db, ID)
	if err != nil {
		return nil, err
	}

	return unit, nil
}

func (u UnitRepository) GetPots(unit milton.Unit) (milton.FlowerPotSlice, error) {
	return unit.UnitFlowerPots().All(context.Background(), u.db)
}

func (u UnitRepository) Pair(name string, mdns string) (milton.Unit, error) {
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

func (u UnitRepository) Unpair(ID string) error {
	unit, err := models.FindUnit(context.Background(), u.db, ID)
	if err != nil {
		return err
	}

	_, err = unit.Delete(context.Background(), u.db)

	return err
}

func (u UnitRepository) All() (milton.UnitSlice, error) {
	units, err := models.Units().All(context.Background(), u.db)

	if err != nil {
		return nil, err
	}

	return units, nil
}

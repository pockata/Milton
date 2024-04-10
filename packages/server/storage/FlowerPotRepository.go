package storage

import (
	"context"
	"database/sql"
	"fmt"
	"milton"
	models "milton/generated_models"

	"github.com/lucsky/cuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type FlowerPotRepository struct {
	db *sql.DB
}

func NewFlowerPotRepository(db *sql.DB) FlowerPotRepository {
	return FlowerPotRepository{
		db: db,
	}
}

func (p FlowerPotRepository) Add(name string, unit milton.Unit) (milton.FlowerPot, error) {
	pot := models.FlowerPot{
		ID:     fmt.Sprintf("fp-%s", cuid.New()),
		Name:   name,
		UnitID: unit.ID,
	}

	if err := pot.Insert(context.Background(), p.db, boil.Infer()); err != nil {
		return nil, err
	}

	return &pot, nil
}

func (p FlowerPotRepository) RemoveByID(ID string) error {
	pot, err := models.FindFlowerPot(context.Background(), p.db, ID)
	if err != nil {
		return err
	}

	return p.Remove(pot)
}

func (p FlowerPotRepository) Remove(pot milton.FlowerPot) error {
	_, err := pot.Delete(context.Background(), p.db)

	return err
}

func (p FlowerPotRepository) All() (milton.FlowerPotSlice, error) {
	pots, err := models.FlowerPots().All(context.Background(), p.db)

	if err != nil {
		return nil, err
	}

	return pots, nil
}

func (p FlowerPotRepository) Get(ID string) (milton.FlowerPot, error) {
	pot, err := models.FindFlowerPot(context.Background(), p.db, ID)

	if err != nil {
		return nil, err
	}

	return pot, err
}

func (p FlowerPotRepository) Update(pot milton.FlowerPot) error {
	_, err := pot.Update(context.Background(), p.db, boil.Infer())

	return err
}

func (p FlowerPotRepository) GetPotsForUnit(unitID string) (milton.FlowerPotSlice, error) {
	mods := []qm.QueryMod{
		models.FlowerPotWhere.UnitID.EQ(unitID),
	}
	pots, err := models.FlowerPots(mods...).All(context.Background(), p.db)

	if err != nil {
		return nil, err
	}

	return pots, nil
}

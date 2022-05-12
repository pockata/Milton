package storage

import (
	"context"
	"database/sql"
	"milton"
	models "milton/generated_models"

	"github.com/lucsky/cuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type FlowerPotService struct {
	db *sql.DB
}

func NewFlowerPotService(db *sql.DB) *FlowerPotService {
	return &FlowerPotService{db: db}
}

func (p *FlowerPotService) Add(name string, unit milton.Unit) (milton.FlowerPot, error) {
	ctx := context.Background()

	pot := models.FlowerPot{
		ID:     cuid.New(),
		Name:   name,
		UnitID: unit.ID(),
	}

	if err := pot.Insert(ctx, p.db, boil.Infer()); err != nil {
		return nil, err
	}

	return transformFlowerPot(&pot), nil
}

func (p *FlowerPotService) Remove(ID string) error {
	ctx := context.Background()

	pot, err := models.FindFlowerPot(ctx, p.db, ID)

	if err != nil {
		return err
	}

	_, err = pot.Delete(ctx, p.db)

	return err
}

func (p *FlowerPotService) All() ([]milton.FlowerPot, error) {
	ctx := context.Background()

	pots, err := models.FlowerPots().All(ctx, p.db)

	if err != nil {
		return nil, err
	}

	mpots := make([]milton.FlowerPot, len(pots))

	for i := range pots {
		mpots[i] = transformFlowerPot(pots[i])
	}

	return mpots, nil
}

func transformFlowerPot(pot *models.FlowerPot) milton.FlowerPot {
	return NewFlowerPot(pot)
}

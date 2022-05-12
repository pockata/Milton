package storage

import (
	models "milton/generated_models"
)

type FlowerPot struct {
	pot *models.FlowerPot
}

func NewFlowerPot(pot *models.FlowerPot) FlowerPot {
	return FlowerPot{
		pot: pot,
	}
}

func (p FlowerPot) ID() string {
	return p.pot.ID
}

func (p FlowerPot) Name() string {
	return p.pot.Name
}

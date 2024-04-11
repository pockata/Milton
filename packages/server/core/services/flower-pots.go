package services

import (
	"fmt"
	"milton/core/domain"
)

func (a App) AddFlowerPot(name string, unitID string) (domain.FlowerPot, error) {
	unit, err := a.unitRepository.Get(unitID)
	if err != nil {
		return nil, fmt.Errorf("couldn't find unit: %w", err)
	}

	flowerPot, err := a.flowerPotRepository.Add(name, unit)
	if err != nil {
		return nil, fmt.Errorf("error inserting flower pot: %w", err)
	}

	return flowerPot, nil
}

func (a App) GetFlowerPots(unitID string) (domain.FlowerPotSlice, error) {
	unit, err := a.unitRepository.Get(unitID)
	if err != nil {
		return nil, fmt.Errorf("couldn't find unit: %w", err)
	}

	pots, err := a.flowerPotRepository.GetPotsForUnit(unit.ID)
	if err != nil {
		return nil, fmt.Errorf("couldn't get flower pots: %w", err)
	}

	return pots, nil
}

func (a App) RenameFlowerPot(potID string, name string) error {
	pot, err := a.flowerPotRepository.Get(potID)
	if err != nil {
		return fmt.Errorf("couldn't find flower pot: %w", err)
	}

	pot.Name = name

	if err := a.flowerPotRepository.Update(pot); err != nil {
		return fmt.Errorf("couldn't update flower pot name: %w", err)
	}

	return nil
}

func (a App) RemoveFlowerPot(potID string) error {
	if err := a.flowerPotRepository.RemoveByID(potID); err != nil {
		return fmt.Errorf("couldn't remove flower pot: %w", err)
	}

	return nil
}

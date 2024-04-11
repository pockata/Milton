package services

import (
	"fmt"
	"milton/core/domain"
	"milton/core/ports"
)

type FlowerPotService struct {
	units      ports.UnitRepository
	flowerPots ports.FlowerPotRepository
}

func NewFlowerPotService(
	units ports.UnitRepository,
	flowerPots ports.FlowerPotRepository,
) FlowerPotService {
	return FlowerPotService{
		units:      units,
		flowerPots: flowerPots,
	}
}

func (s FlowerPotService) AddFlowerPot(name string, unitID string) (domain.FlowerPot, error) {
	unit, err := s.units.Get(unitID)
	if err != nil {
		return nil, fmt.Errorf("couldn't find unit: %w", err)
	}

	flowerPot, err := s.flowerPots.Add(name, unit)
	if err != nil {
		return nil, fmt.Errorf("error inserting flower pot: %w", err)
	}

	return flowerPot, nil
}

func (s FlowerPotService) GetFlowerPots(unitID string) (domain.FlowerPotSlice, error) {
	unit, err := s.units.Get(unitID)
	if err != nil {
		return nil, fmt.Errorf("couldn't find unit: %w", err)
	}

	pots, err := s.flowerPots.GetPotsForUnit(unit.ID)
	if err != nil {
		return nil, fmt.Errorf("couldn't get flower pots: %w", err)
	}

	return pots, nil
}

func (s FlowerPotService) RenameFlowerPot(potID string, name string) error {
	pot, err := s.flowerPots.Get(potID)
	if err != nil {
		return fmt.Errorf("couldn't find flower pot: %w", err)
	}

	pot.Name = name

	if err := s.flowerPots.Update(pot); err != nil {
		return fmt.Errorf("couldn't update flower pot name: %w", err)
	}

	return nil
}

func (s FlowerPotService) RemoveFlowerPot(potID string) error {
	if err := s.flowerPots.RemoveByID(potID); err != nil {
		return fmt.Errorf("couldn't remove flower pot: %w", err)
	}

	return nil
}

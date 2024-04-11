package services

import (
	"fmt"
	"milton/core/domain"
	"milton/core/ports"
)

type UnitService struct {
	units      ports.UnitRepository
	flowerPots ports.FlowerPotRepository
}

func NewUnitService(
	units ports.UnitRepository,
	flowerPots ports.FlowerPotRepository,
) UnitService {
	return UnitService{
		units:      units,
		flowerPots: flowerPots,
	}
}

func (s UnitService) Pair(name string, mdns string) (domain.Unit, error) {
	unit, err := s.units.Pair(mdns, name)
	if err != nil {
		return nil, fmt.Errorf("couldn't pair unit: %w", err)
	}

	return unit, nil
}

func (s UnitService) GetAll() (domain.UnitSlice, error) {
	units, err := s.units.All()
	if err != nil {
		return nil, fmt.Errorf("couldn't get all units: %w", err)
	}

	return units, nil
}

func (s UnitService) Unpair(ID string) error {
	unit, err := s.units.Get(ID)
	if err != nil {
		return fmt.Errorf("couldn't find unit to delete: %w", err)
	}

	pots, err := s.flowerPots.GetPotsForUnit(unit.ID)
	if err != nil {
		return fmt.Errorf("couldn't get unit pots: %w", err)
	}

	for _, pot := range pots {
		if err := s.flowerPots.Remove(pot); err != nil {
			return fmt.Errorf("couldn't remove flower pot from unit: %w", err)
		}
	}

	if err := s.units.Unpair(unit.ID); err != nil {
		return fmt.Errorf("couldn't unpair unit: %w", err)
	}

	return nil
}

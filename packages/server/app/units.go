package app

import (
	"fmt"
	"milton"
)

func (a App) PairUnit(name string, mdns string) (milton.Unit, error) {
	unit, err := a.unitRepository.Pair(mdns, name)
	if err != nil {
		return nil, fmt.Errorf("couldn't pair unit: %w", err)
	}

	return unit, nil
}

func (a App) GetAllUnits() (milton.UnitSlice, error) {
	units, err := a.unitRepository.All()
	if err != nil {
		return nil, fmt.Errorf("couldn't get all units: %w", err)
	}

	return units, nil
}

func (a App) UnpairUnit(ID string) error {
	unit, err := a.unitRepository.Get(ID)
	if err != nil {
		return fmt.Errorf("couldn't find unit to delete: %w", err)
	}

	pots, err := a.flowerPotRepository.GetPotsForUnit(unit.ID)
	if err != nil {
		return fmt.Errorf("couldn't get unit pots: %w", err)
	}

	for _, pot := range pots {
		if err := a.flowerPotRepository.Remove(pot); err != nil {
			return fmt.Errorf("couldn't remove flower pot from unit: %w", err)
		}
	}

	if err := a.unitRepository.Unpair(unit.ID); err != nil {
		return fmt.Errorf("couldn't unpair unit: %w", err)
	}

	return nil
}

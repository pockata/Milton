package services

import (
	"milton/core/ports"
)

type App struct {
	flowerPots ports.FlowerPotRepository
	units      ports.UnitRepository
	jobs       ports.JobRepository
}

type AppConfig struct {
	FlowerPots ports.FlowerPotRepository
	Units      ports.UnitRepository
	Jobs       ports.JobRepository
}

func NewApp(cfg AppConfig) App {
	return App{
		flowerPots: cfg.FlowerPots,
		units:      cfg.Units,
		jobs:       cfg.Jobs,
	}
}

package app

import "milton"

type App struct {
	flowerPotService milton.FlowerPotService
	unitService      milton.UnitService
	jobService       milton.JobService
}

type AppConfig struct {
	FlowerPotService milton.FlowerPotService
	UnitService      milton.UnitService
	JobService       milton.JobService
}

func NewApp(cfg AppConfig) App {
	return App{
		flowerPotService: cfg.FlowerPotService,
		unitService:      cfg.UnitService,
		jobService:       cfg.JobService,
	}
}

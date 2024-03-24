package app

import (
	"milton/storage"
)

type App struct {
	flowerPotService storage.FlowerPotService
	unitService      storage.UnitService
	jobService       storage.JobService
}

type AppConfig struct {
	FlowerPotService storage.FlowerPotService
	UnitService      storage.UnitService
	JobService       storage.JobService
}

func NewApp(cfg AppConfig) App {
	return App{
		flowerPotService: cfg.FlowerPotService,
		unitService:      cfg.UnitService,
		jobService:       cfg.JobService,
	}
}

package app

import (
	"milton/storage"
)

type App struct {
	flowerPotService storage.FlowerPotService
	unitService      storage.UnitService
}

type AppConfig struct {
	FlowerPotService storage.FlowerPotService
	UnitService      storage.UnitService
}

func NewApp(cfg AppConfig) App {
	return App{
		flowerPotService: cfg.FlowerPotService,
		unitService:      cfg.UnitService,
	}
}

package app

import "milton"

type App struct {
	flowerPotRepository milton.FlowerPotRepository
	unitRepository      milton.UnitRepository
	jobRepository       milton.JobRepository
}

type AppConfig struct {
	FlowerPotRepository milton.FlowerPotRepository
	UnitRepository      milton.UnitRepository
	JobRepository       milton.JobRepository
}

func NewApp(cfg AppConfig) App {
	return App{
		flowerPotRepository: cfg.FlowerPotRepository,
		unitRepository:      cfg.UnitRepository,
		jobRepository:       cfg.JobRepository,
	}
}

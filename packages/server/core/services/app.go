package services

import (
	"milton/core/ports"
)

type App struct {
	flowerPotRepository ports.FlowerPotRepository
	unitRepository      ports.UnitRepository
	jobRepository       ports.JobRepository
}

type AppConfig struct {
	FlowerPotRepository ports.FlowerPotRepository
	UnitRepository      ports.UnitRepository
	JobRepository       ports.JobRepository
}

func NewApp(cfg AppConfig) App {
	return App{
		flowerPotRepository: cfg.FlowerPotRepository,
		unitRepository:      cfg.UnitRepository,
		jobRepository:       cfg.JobRepository,
	}
}

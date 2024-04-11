package ports

import "milton/core/domain"

type FlowerPotService interface {
	AddFlowerPot(string, string) (domain.FlowerPot, error)
	GetFlowerPots(string) (domain.FlowerPotSlice, error)
	RenameFlowerPot(string, string) error
	RemoveFlowerPot(string) error
}

type JobService interface {
	AddJob(JobCreateConfig) (domain.Job, error)
	RemoveJob(string) error
	UpdateJob(string, JobUpdateConfig) (domain.Job, error)
	GetJob(string) (domain.Job, error)
	GetAllJobs() (domain.JobSlice, error)
}

type UnitService interface {
	PairUnit(string, string) (domain.Unit, error)
	GetAllUnits() (domain.UnitSlice, error)
	UnpairUnit(string) error
}

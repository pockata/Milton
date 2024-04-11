package ports

import "milton/core/domain"

type FlowerPotService interface {
	Add(string, string) (domain.FlowerPot, error)
	GetAll(string) (domain.FlowerPotSlice, error)
	Rename(string, string) error
	Remove(string) error
}

type JobService interface {
	Add(JobCreateConfig) (domain.Job, error)
	Remove(string) error
	Update(string, JobUpdateConfig) (domain.Job, error)
	Get(string) (domain.Job, error)
	GetAll() (domain.JobSlice, error)
}

type UnitService interface {
	Pair(string, string) (domain.Unit, error)
	GetAll() (domain.UnitSlice, error)
	Unpair(string) error
}

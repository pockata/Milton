package ports

import (
	"milton/core/domain"
	"time"
)

type Logger interface {
	Info(string, ...any)
	Error(string, ...any)
}

type UnitRepository interface {
	Pair(string, string) (domain.Unit, error)
	Unpair(string) error
	All() (domain.UnitSlice, error)
	Get(string) (domain.Unit, error)
}

type FlowerPotRepository interface {
	Add(string, domain.Unit) (domain.FlowerPot, error)
	RemoveByID(string) error
	Remove(domain.FlowerPot) error
	Get(string) (domain.FlowerPot, error)
	GetPotsForUnit(string) (domain.FlowerPotSlice, error)
	All() (domain.FlowerPotSlice, error)
	Update(domain.FlowerPot) error
}

type JobRepository interface {
	Get(string) (domain.Job, error)
	GetAll() (domain.JobSlice, error)
	Remove(string) error
	Add(JobCreateConfig) (domain.Job, error)
	Update(string, JobUpdateConfig) (domain.Job, error)
}

type JobCreateConfig struct {
	UnitID      string
	FlowerPotID string
	StartTime   time.Time
	WaterQty    int64
}

type JobUpdateConfig struct {
	StartTime *time.Time
	Status    *domain.JobStatus
	WaterQty  *int64
}

type LogRepository interface {
	Add(LogCreateConfig) error
	GetAll() (string, error)
}

type LogCreateConfig struct {
	Unit    domain.Unit
	Job     domain.Job
	Message string
}

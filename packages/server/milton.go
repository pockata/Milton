package milton

import "time"

type UnitService interface {
	Pair(string, string) error
	Unpair(string) error
	All() ([]Unit, error)
}

type Unit interface {
	ID() string
	Name() string
	MDNS() string
}

type FlowerPotService interface {
	Add(string, Unit) (FlowerPot, error)
	Remove(string) error
	All() ([]FlowerPot, error)
}

type FlowerPot interface {
	ID() string
	Name() string
	Update(name string) error
}

type JobStatus uint8

const (
	Pending JobStatus = iota
	Running
	Complete
	Error
)

type Job interface {
	Unit() (Unit, error)
	FlowerPot() (FlowerPot, error)
	StartTime() time.Time
	WaterQty() int64
	Status() JobStatus
	Remove() error
}

type JobService interface {
	Get(string) (Job, error)
	GetAll() ([]Job, error)
	Remove(string) error
	Add(JobCreateConfig) (Job, error)
}

type JobCreateConfig struct {
	Unit      Unit
	FlowerPot FlowerPot
	StartTime time.Time
	WaterQty  int64
	Status    JobStatus
}

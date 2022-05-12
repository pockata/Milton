package milton

import "time"

type UnitService interface {
	Pair(string, string) error
	Unpair(string) error
	All() ([]Unit, error)
	Get(string) (Unit, error)
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
}

type JobStatus uint8

const (
	Pending JobStatus = iota + 1
	Running
	Complete
	Error
)

type Job interface {
	ID() string
	Unit() Unit
	FlowerPot() FlowerPot
	StartTime() time.Time
	WaterQty() int64
	Status() JobStatus
}

type JobUpdateConfig struct {
	StartTime *time.Time
	Status    *JobStatus
	WaterQty  *int64
}

type JobService interface {
	Get(string) (Job, error)
	GetAll() ([]Job, error)
	Remove(string) error
	Add(JobCreateConfig) (Job, error)
	Update(JobUpdateConfig) (Job, error)
}

type JobCreateConfig struct {
	Unit      Unit
	FlowerPot FlowerPot
	StartTime time.Time
	WaterQty  int64
	Status    JobStatus
}

type LogService interface {
	Add(LogCreateConfig) error
	GetAll() (string, error)
}

type LogCreateConfig struct {
	Unit    Unit
	Job     Job
	Message string
}

type LogFormat struct {
	Message   string       `json:"message"`
	Unit      LogUnit      `json:"unit"`
	Job       LogJob       `json:"job"`
	FlowerPot LogFlowerPot `json:"flowerPot"`
}

type LogUnit struct {
	Name string `json:"name"`
	MDNS string `json:"mdns"`
}

type LogJob struct {
	StartTime time.Time `json:"startTime"`
	WaterQty  int64     `json:"waterQty"`
	Status    JobStatus `json:"status"`
}

type LogFlowerPot struct {
	Name string `json:"name"`
}

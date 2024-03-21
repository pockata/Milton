package milton

import (
	models "milton/generated_models"
	"time"
)

// Build is the git version of this program. It is set using Build flags in the
// Makefile.
var Build = "develop"

type Logger interface {
	Info(string, ...any)
	Error(string, ...any)
}

type UnitService interface {
	Pair(string, string) error
	Unpair(string) error
	All() (UnitSlice, error)
	Get(string) (Unit, error)
}

type Unit *models.Unit
type UnitSlice models.UnitSlice

type FlowerPotService interface {
	Add(string, Unit) (FlowerPot, error)
	Remove(string) error
	All() (FlowerPotSlice, error)
}

type FlowerPot *models.FlowerPot
type FlowerPotSlice models.FlowerPotSlice

type JobStatus uint8

const (
	Pending JobStatus = iota + 1
	Running
	Complete
	Error
)

type Job *models.Job
type JobSlice models.JobSlice

type JobUpdateConfig struct {
	StartTime *time.Time
	Status    *JobStatus
	WaterQty  *int64
}

type JobService interface {
	Get(string) (Job, error)
	GetAll() (JobSlice, error)
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

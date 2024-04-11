package domain

import (
	models "milton/adapters/db/generated_models"
	"time"
)

type Unit = *models.Unit
type UnitSlice = models.UnitSlice

type FlowerPot = *models.FlowerPot
type FlowerPotSlice = models.FlowerPotSlice

type Job = *models.Job
type JobSlice = models.JobSlice

type JobStatus uint8

const (
	Pending JobStatus = iota + 1
	Running
	Complete
	Error
)

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

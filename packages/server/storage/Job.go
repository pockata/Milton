package storage

import (
	"milton"
	models "milton/generated_models"
	"time"
)

type Job struct {
	job *models.Job
}

func NewJob(job *models.Job) Job {
	return Job{
		job: job,
	}
}

func (j Job) ID() string {
	return j.job.ID
}

func (j Job) Unit() milton.Unit {
	return transformUnit(j.job.R.Unit)
}

func (j Job) FlowerPot() milton.FlowerPot {
	return transformFlowerPot(j.job.R.FlowerPot)
}

func (j Job) StartTime() time.Time {
	return j.job.StartTime
}

func (j Job) WaterQty() int64 {
	return j.job.WaterQty
}

func (j Job) Status() milton.JobStatus {
	return milton.JobStatus(j.job.Status)
}

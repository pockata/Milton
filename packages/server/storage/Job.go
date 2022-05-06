package storage

import (
	"context"
	"database/sql"
	"milton"
	models "milton/generated_models"
	"time"
)

type Job struct {
	db  *sql.DB
	job *models.Job
}

func (j *Job) Unit() (milton.Unit, error) {
	ctx := context.Background()

	unit, err := j.job.Unit().One(ctx, j.db)
	if err != nil {
		return nil, err
	}

	return transformUnit(unit, j.db), nil
}

func (j *Job) FlowerPot() (milton.FlowerPot, error) {
	ctx := context.Background()

	pot, err := j.job.FlowerPot().One(ctx, j.db)
	if err != nil {
		return nil, err
	}

	return transformFlowerPot(pot, j.db), nil
}

func (j *Job) StartTime() time.Time {
	return j.job.StartTime
}

func (j *Job) WaterQty() int64 {
	return j.job.WaterQty
}

func (j *Job) Status() milton.JobStatus {
	return milton.JobStatus(j.job.Status)
}

func (j *Job) Remove() error {
	ctx := context.Background()
	_, err := j.job.Delete(ctx, j.db)

	return err
}

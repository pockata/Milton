package storage

import (
	"context"
	"database/sql"
	"milton"
	models "milton/generated_models"

	"github.com/lucsky/cuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type JobService struct {
	db *sql.DB
}

func (s *JobService) Get(ID string) (milton.Job, error) {
	ctx := context.Background()
	job, err := models.FindJob(ctx, s.db, ID)

	if err != nil {
		return nil, err
	}

	return transformJob(job, s.db), nil
}

func (s *JobService) GetAll() ([]milton.Job, error) {
	ctx := context.Background()
	jobs, err := models.Jobs().All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	tjobs := make([]milton.Job, len(jobs))

	for i, jb := range jobs {
		tjobs[i] = transformJob(jb, s.db)
	}

	return tjobs, err
}

func (s *JobService) Remove(ID string) error {
	ctx := context.Background()
	job, err := models.FindJob(ctx, s.db, ID, models.JobColumns.ID)

	if err != nil {
		return err
	}

	_, err = job.Delete(ctx, s.db)

	return err
}

func (s *JobService) Add(cfg milton.JobCreateConfig) (milton.Job, error) {
	ctx := context.Background()

	job := &models.Job{
		ID:          cuid.New(),
		StartTime:   cfg.StartTime,
		WaterQty:    cfg.WaterQty,
		Status:      int64(cfg.Status),
		UnitID:      cfg.Unit.ID(),
		FlowerPotID: cfg.FlowerPot.ID(),
	}

	if err := job.Insert(ctx, s.db, boil.Infer()); err != nil {
		return nil, err
	}

	return transformJob(job, s.db), nil
}

func transformJob(job *models.Job, db *sql.DB) milton.Job {
	return &Job{
		job: job,
		db:  db,
	}
}

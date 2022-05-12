package storage

import (
	"context"
	"database/sql"
	"errors"
	"milton"
	models "milton/generated_models"

	"github.com/lucsky/cuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type JobService struct {
	db *sql.DB
}

func NewJobService(db *sql.DB) *JobService {
	return &JobService{db: db}
}

func (s *JobService) Get(ID string) (milton.Job, error) {
	ctx := context.Background()
	job, err := models.FindJob(ctx, s.db, ID)

	if err != nil {
		return nil, err
	}

	// load the relationships
	job.Unit().One(ctx, s.db)
	job.FlowerPot().One(ctx, s.db)

	return transformJob(job), nil
}

func (s *JobService) GetAll() ([]milton.Job, error) {
	ctx := context.Background()
	potRel := qm.Load(models.JobRels.FlowerPot)
	unitRel := qm.Load(models.JobRels.Unit)
	jobs, err := models.Jobs(potRel, unitRel).All(ctx, s.db)

	if err != nil {
		return nil, err
	}

	tjobs := make([]milton.Job, len(jobs))

	for i, jb := range jobs {
		tjobs[i] = transformJob(jb)
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

	return transformJob(job), nil
}

func (s *JobService) Update(ID string, upd milton.JobUpdateConfig) (milton.Job, error) {
	ctx := context.Background()
	job, err := models.FindJob(ctx, s.db, ID)

	if err != nil {
		return nil, err
	}

	// not too happy with this approach

	// TODO: Test if boil.Infer() doesn't prevent updating in this case
	hasChangedField := false

	if upd.StartTime != nil {
		job.StartTime = *upd.StartTime
		hasChangedField = true
	}

	if upd.WaterQty != nil {
		job.WaterQty = *upd.WaterQty
		hasChangedField = true
	}

	if upd.Status != nil {
		job.Status = int64(*upd.Status)
		hasChangedField = true
	}

	if !hasChangedField {
		return nil, errors.New("No values passed for updating")
	}

	_, err = job.Update(ctx, s.db, boil.Infer())

	return transformJob(job), nil
}

func transformJob(job *models.Job) milton.Job {
	return Job{
		job: job,
	}
}

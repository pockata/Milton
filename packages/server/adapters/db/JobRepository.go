package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	models "milton/adapters/db/generated_models"
	"milton/core/domain"
	"milton/core/ports"

	"github.com/lucsky/cuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type JobRepository struct {
	db *sql.DB
}

func NewJobRepository(db *sql.DB) JobRepository {
	return JobRepository{
		db: db,
	}
}

func (s JobRepository) Get(ID string) (domain.Job, error) {
	ctx := context.Background()
	mods := []qm.QueryMod{
		models.JobWhere.ID.EQ(ID),
		qm.Load(models.JobRels.Unit),
		qm.Load(models.JobRels.FlowerPot),
	}

	return models.Jobs(mods...).One(ctx, s.db)
}

func (s JobRepository) GetAll() (domain.JobSlice, error) {
	ctx := context.Background()
	mods := []qm.QueryMod{
		qm.Load(models.JobRels.FlowerPot),
		qm.Load(models.JobRels.Unit),
	}

	return models.Jobs(mods...).All(ctx, s.db)
}

func (s JobRepository) Remove(ID string) error {
	ctx := context.Background()
	job, err := models.FindJob(ctx, s.db, ID, models.JobColumns.ID)

	if err != nil {
		return err
	}

	_, err = job.Delete(ctx, s.db)

	return err
}

func (s JobRepository) Add(cfg ports.JobCreateConfig) (domain.Job, error) {
	job := &models.Job{
		ID:          fmt.Sprintf("j-%s", cuid.New()),
		StartTime:   cfg.StartTime,
		WaterQty:    cfg.WaterQty,
		UnitID:      cfg.UnitID,
		FlowerPotID: cfg.FlowerPotID,
	}

	if err := job.Insert(context.Background(), s.db, boil.Infer()); err != nil {
		return nil, err
	}

	return job, nil
}

func (s JobRepository) Update(ID string, upd ports.JobUpdateConfig) (domain.Job, error) {
	ctx := context.Background()
	job, err := models.FindJob(ctx, s.db, ID)

	if err != nil {
		return nil, err
	}

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
		return nil, errors.New("no values passed for updating")
	}

	_, err = job.Update(ctx, s.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	return job, nil
}

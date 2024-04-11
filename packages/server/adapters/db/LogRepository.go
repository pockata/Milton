package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	models "milton/adapters/db/generated_models"
	"milton/core/domain"
	"milton/core/ports"

	"github.com/lucsky/cuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type LogRepository struct {
	db *sql.DB
}

func NewLogRepository(db *sql.DB) *LogRepository {
	return &LogRepository{
		db: db,
	}
}

func (l *LogRepository) Add(cfg ports.LogCreateConfig) error {
	ctx := context.Background()

	log := &models.Log{
		ID:      fmt.Sprintf("l-%s", cuid.New()),
		UnitID:  cfg.Unit.ID,
		JobID:   cfg.Job.ID,
		Message: cfg.Message,
	}

	return log.Insert(ctx, l.db, boil.Infer())
}

func (l *LogRepository) GetAll() (string, error) {
	ctx := context.Background()

	logs, err := models.Logs(
		qm.Load(models.LogRels.Unit),
		qm.Load(models.LogRels.Job, qm.Load(models.JobRels.FlowerPot)),
	).All(ctx, l.db)

	if err != nil {
		return "", err
	}

	rows := make([]domain.LogFormat, len(logs))

	for i, log := range logs {
		rows[i] = domain.LogFormat{
			Message: log.Message,
			Unit: domain.LogUnit{
				Name: log.R.Unit.Name,
				MDNS: log.R.Unit.MDNS,
			},
			Job: domain.LogJob{
				StartTime: log.R.Job.StartTime,
				Status:    domain.JobStatus(log.R.Job.Status),
				WaterQty:  log.R.Job.WaterQty,
			},
			FlowerPot: domain.LogFlowerPot{
				Name: log.R.Job.R.FlowerPot.Name,
			},
		}
	}

	jsonData, err := json.Marshal(rows)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

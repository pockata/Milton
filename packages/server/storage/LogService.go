package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"milton"
	models "milton/generated_models"

	"github.com/lucsky/cuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type LogService struct {
	db *sql.DB
}

func NewLogService(db *sql.DB) *LogService {
	return &LogService{db: db}
}

func (l *LogService) Add(cfg milton.LogCreateConfig) error {
	ctx := context.Background()

	log := &models.Log{
		ID:      cuid.New(),
		UnitID:  cfg.Unit.ID(),
		JobID:   cfg.Job.ID(),
		Message: cfg.Message,
	}

	return log.Insert(ctx, l.db, boil.Infer())
}

func (l *LogService) GetAll() (string, error) {
	ctx := context.Background()

	logs, err := models.Logs(
		qm.Load(models.LogRels.Unit),
		qm.Load(models.LogRels.Job, qm.Load(models.JobRels.FlowerPot)),
	).All(ctx, l.db)

	if err != nil {
		return "", err
	}

	rows := make([]milton.LogFormat, len(logs))

	for i, log := range logs {
		rows[i] = milton.LogFormat{
			Message: log.Message,
			Unit: milton.LogUnit{
				Name: log.R.Unit.Name,
				MDNS: log.R.Unit.MDNS,
			},
			Job: milton.LogJob{
				StartTime: log.R.Job.StartTime,
				Status:    milton.JobStatus(log.R.Job.Status),
				WaterQty:  log.R.Job.WaterQty,
			},
			FlowerPot: milton.LogFlowerPot{
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

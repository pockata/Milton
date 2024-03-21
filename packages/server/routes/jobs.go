package routes

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/lucsky/cuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "milton/generated_models"
	"milton/helpers"
)

type AddJobResponse struct {
	Job models.Job `json:"job"`
}

func AddJob(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if err := r.ParseForm(); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	unitID := r.PostForm.Get("UnitID")
	potID := r.PostForm.Get("PotID")
	startTimeStr := r.PostForm.Get("StartTime")
	statusStr := r.PostForm.Get("Status")
	waterQtyStr := r.PostForm.Get("WaterQty")

	if !helpers.CheckParams(unitID, potID, waterQtyStr, startTimeStr, statusStr) {
		helpers.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	startTimeInt, _ := strconv.ParseInt(startTimeStr, 10, 64)
	startTime := time.Unix(startTimeInt, 0)
	waterQty, _ := strconv.Atoi(waterQtyStr)
	status, _ := strconv.Atoi(statusStr)

	if startTime.Before(time.Now()) {
		helpers.ErrorResponse(w, r, errors.New("start time should be in the future"))
		return
	}

	job := models.Job{
		ID:          fmt.Sprintf("j-%s", cuid.New()),
		UnitID:      unitID,
		FlowerPotID: potID,
		WaterQty:    int64(waterQty),
		StartTime:   startTime,
		Status:      int64(status),
	}

	if err := job.Insert(context.Background(), db, boil.Infer()); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't create job: %w", err))
		return
	}

	helpers.SuccessResponse(w, r, AddJobResponse{
		Job: job,
	})
}

type RemoveJobResponse struct {
	Success bool `json:"success"`
}

func RemoveJob(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if err := r.ParseForm(); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	ID := r.Form.Get("ID")
	if !helpers.CheckParams(ID) {
		helpers.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	_, err := models.FindJob(context.Background(), db, ID)
	if err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't find job: %w", err))
		return
	}

	helpers.SuccessResponse(w, r, RemoveJobResponse{
		Success: true,
	})
}

type GetJobResponse struct {
	Job models.Job `json:"job"`
}

func GetJob(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	jobID := r.PathValue("JobID")

	if !helpers.CheckParams(jobID) {
		helpers.ErrorResponse(w, r, errors.New("invalid job ID"))
		return
	}

	mods := []qm.QueryMod{
		models.JobWhere.ID.EQ(jobID),
		qm.Load(models.JobRels.FlowerPot),
		qm.Load(models.JobRels.Unit),
	}

	job, err := models.Jobs(mods...).One(context.Background(), db)
	if err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't get job: %w", err))
		return
	}

	helpers.SuccessResponse(w, r, GetJobResponse{
		Job: *job,
	})
}

type GetJobsResponse struct {
	Jobs models.JobSlice `json:"jobs"`
}

func GetJobs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	mods := []qm.QueryMod{
		qm.Load(models.JobRels.FlowerPot),
		qm.Load(models.JobRels.Unit),
	}

	jobs, err := models.Jobs(mods...).All(context.Background(), db)
	if err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't get jobs: %w", err))
		return
	}

	helpers.SuccessResponse(w, r, GetJobsResponse{
		Jobs: jobs,
	})
}

type UpdateJobResponse struct {
	Job models.Job `json:"job"`
}

func UpdateJob(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if err := r.ParseForm(); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	jobID := r.PostForm.Get("JobID")
	startTimeStr := r.PostForm.Get("StartTime")
	statusStr := r.PostForm.Get("Status")
	waterQtyStr := r.PostForm.Get("WaterQty")

	if !helpers.CheckParams(jobID, waterQtyStr, startTimeStr, statusStr) {
		helpers.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	startTimeInt, _ := strconv.ParseInt(startTimeStr, 10, 64)
	startTime := time.Unix(startTimeInt, 0)
	waterQty, _ := strconv.Atoi(waterQtyStr)
	status, _ := strconv.Atoi(statusStr)

	if startTime.Before(time.Now()) {
		helpers.ErrorResponse(w, r, errors.New("start time should be in the future"))
		return
	}

	job, err := models.FindJob(context.Background(), db, jobID)
	if err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't get job: %w", err))
		return
	}

	job.WaterQty = int64(waterQty)
	job.StartTime = startTime
	job.Status = int64(status)

	if _, err := job.Update(context.Background(), db, boil.Infer()); err != nil {
		helpers.ErrorResponse(w, r, fmt.Errorf("couldn't update job: %w", err))
		return
	}

	helpers.SuccessResponse(w, r, UpdateJobResponse{
		Job: *job,
	})
}

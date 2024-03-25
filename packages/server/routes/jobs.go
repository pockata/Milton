package routes

import (
	"errors"
	"fmt"
	"milton"
	"net/http"
	"strconv"
	"time"

	"milton/helpers"
)

type AddJobResponse struct {
	Job milton.Job `json:"job"`
}

func (c Controller) AddJob(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	unitID := r.PostForm.Get("UnitID")
	potID := r.PostForm.Get("PotID")
	startTimeStr := r.PostForm.Get("StartTime")
	waterQtyStr := r.PostForm.Get("WaterQty")

	if !helpers.ValidParams(unitID, potID, waterQtyStr, startTimeStr) {
		c.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	startTimeInt, err := strconv.ParseInt(startTimeStr, 10, 64)
	if err != nil {
		c.ErrorResponse(w, r, errors.New("invalid start time"))
		return
	}

	startTime := time.Unix(startTimeInt, 0)

	waterQty, err := strconv.Atoi(waterQtyStr)
	if err != nil {
		c.ErrorResponse(w, r, errors.New("invalid water quantity"))
		return
	}

	if startTime.Before(time.Now()) {
		c.ErrorResponse(w, r, errors.New("start time should be in the future"))
		return
	}

	job, err := c.app.AddJob(milton.JobCreateConfig{
		UnitID:      unitID,
		FlowerPotID: potID,
		StartTime:   startTime,
		WaterQty:    int64(waterQty),
	})
	if err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("couldn't create job: %w", err))
		return
	}

	c.SuccessResponse(w, r, AddJobResponse{
		Job: job,
	})
}

type RemoveJobResponse struct {
	Success bool `json:"success"`
}

func (c Controller) RemoveJob(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	ID := r.Form.Get("ID")
	if !helpers.ValidParams(ID) {
		c.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	if err := c.app.RemoveJob(ID); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.SuccessResponse(w, r, RemoveJobResponse{
		Success: true,
	})
}

type GetJobResponse struct {
	Job milton.Job `json:"job"`
}

func (c Controller) GetJob(w http.ResponseWriter, r *http.Request) {
	jobID := r.PathValue("JobID")

	if !helpers.ValidParams(jobID) {
		c.ErrorResponse(w, r, errors.New("invalid job ID"))
		return
	}

	job, err := c.app.GetJob(jobID)
	if err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("couldn't get job: %w", err))
		return
	}

	c.SuccessResponse(w, r, GetJobResponse{
		Job: job,
	})
}

type GetJobsResponse struct {
	Jobs milton.JobSlice `json:"jobs"`
}

func (c Controller) GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := c.app.GetAllJobs()
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.SuccessResponse(w, r, GetJobsResponse{
		Jobs: jobs,
	})
}

type UpdateJobResponse struct {
	Job milton.Job `json:"job"`
}

func (c Controller) UpdateJob(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	jobID := r.PostForm.Get("JobID")
	startTimeStr := r.PostForm.Get("StartTime")
	statusStr := r.PostForm.Get("Status")
	waterQtyStr := r.PostForm.Get("WaterQty")

	if !helpers.ValidParams(jobID, waterQtyStr, startTimeStr, statusStr) {
		c.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	startTimeInt, err := strconv.ParseInt(startTimeStr, 10, 64)
	if err != nil {
		c.ErrorResponse(w, r, errors.New("invalid start time"))
		return
	}

	startTime := time.Unix(startTimeInt, 0)

	waterQty, err := strconv.Atoi(waterQtyStr)
	if err != nil {
		c.ErrorResponse(w, r, errors.New("invalid water quantity"))
		return
	}

	status, err := strconv.Atoi(statusStr)
	if err != nil {
		c.ErrorResponse(w, r, errors.New("invalid status"))
		return
	}

	if startTime.Before(time.Now()) {
		c.ErrorResponse(w, r, errors.New("start time should be in the future"))
		return
	}

	jStatus := milton.JobStatus(status)
	jWater := int64(waterQty)

	job, err := c.app.UpdateJob(jobID, milton.JobUpdateConfig{
		StartTime: &startTime,
		Status:    &jStatus,
		WaterQty:  &jWater,
	})

	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.SuccessResponse(w, r, UpdateJobResponse{
		Job: job,
	})
}

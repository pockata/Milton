package http

import (
	"errors"
	"fmt"
	"milton/core/domain"
	"milton/core/ports"
	"net/http"
	"strconv"
	"time"
)

type AddJobResponse struct {
	Job domain.Job `json:"job"`
}

func (c HTTPController) AddJob(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	unitID := r.PostForm.Get("UnitID")
	potID := r.PostForm.Get("PotID")
	startTimeStr := r.PostForm.Get("StartTime")
	waterQtyStr := r.PostForm.Get("WaterQty")

	if !c.ValidParams(unitID, potID, waterQtyStr, startTimeStr) {
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

	job, err := c.jobs.AddJob(ports.JobCreateConfig{
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

func (c HTTPController) RemoveJob(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	ID := r.Form.Get("ID")
	if !c.ValidParams(ID) {
		c.ErrorResponse(w, r, errors.New("invalid request. missing parameters"))
		return
	}

	if err := c.jobs.RemoveJob(ID); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.SuccessResponse(w, r, RemoveJobResponse{
		Success: true,
	})
}

type GetJobResponse struct {
	Job domain.Job `json:"job"`
}

func (c HTTPController) GetJob(w http.ResponseWriter, r *http.Request) {
	jobID := r.PathValue("JobID")

	if !c.ValidParams(jobID) {
		c.ErrorResponse(w, r, errors.New("invalid job ID"))
		return
	}

	job, err := c.jobs.GetJob(jobID)
	if err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("couldn't get job: %w", err))
		return
	}

	c.SuccessResponse(w, r, GetJobResponse{
		Job: job,
	})
}

type GetJobsResponse struct {
	Jobs domain.JobSlice `json:"jobs"`
}

func (c HTTPController) GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := c.jobs.GetAllJobs()
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.SuccessResponse(w, r, GetJobsResponse{
		Jobs: jobs,
	})
}

type UpdateJobResponse struct {
	Job domain.Job `json:"job"`
}

func (c HTTPController) UpdateJob(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.ErrorResponse(w, r, fmt.Errorf("error parsing form data: %w", err))
		return
	}

	jobID := r.PostForm.Get("JobID")
	startTimeStr := r.PostForm.Get("StartTime")
	statusStr := r.PostForm.Get("Status")
	waterQtyStr := r.PostForm.Get("WaterQty")

	if !c.ValidParams(jobID, waterQtyStr, startTimeStr, statusStr) {
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

	jStatus := domain.JobStatus(status)
	jWater := int64(waterQty)

	job, err := c.jobs.UpdateJob(jobID, ports.JobUpdateConfig{
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

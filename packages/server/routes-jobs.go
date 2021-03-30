package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"milton/helpers"
	"milton/models"
)

func addJob(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data", err)
		return
	}

	unitID := r.PostForm.Get("UnitID")
	potID := r.PostForm.Get("PotID")
	startTimeStr := r.PostForm.Get("StartTime")
	statusStr := r.PostForm.Get("Status")
	waterQtyStr := r.PostForm.Get("WaterQty")

	if !helpers.CheckParams(unitID, potID, waterQtyStr, startTimeStr, statusStr) {
		helpers.ErrorResponse(rw, errors.New("Invalid request. Missing parameters"))
		return
	}

	startTimeInt, _ := strconv.ParseInt(startTimeStr, 10, 64)
	startTime := time.Unix(startTimeInt, 0)
	waterQty, _ := strconv.Atoi(waterQtyStr)
	status, _ := strconv.Atoi(statusStr)

	if startTime.Before(time.Now()) {
		helpers.ErrorResponse(rw, errors.New("Start time should be in the future"))
		return
	}

	var unit models.Unit
	var pot models.Pot

	findUnit := db.Instance.First(&unit, unitID)
	if findUnit.Error != nil {
		helpers.ErrorResponse(rw, findUnit.Error)
		return
	}

	findPot := db.Instance.First(&pot, potID)
	if findPot.Error != nil {
		helpers.ErrorResponse(rw, findPot.Error)
		return
	}

	entry := &models.Job{
		Unit:      unit,
		Pot:       pot,
		WaterQty:  waterQty,
		StartTime: startTime,
		Status:    status,
	}

	helpers.CreateEntry(rw, r, *db.Instance, &entry)
}

func removeJob(rw http.ResponseWriter, r *http.Request) {
	helpers.DeleteEntry(rw, r, *db.Instance, &models.Job{})
}

func getJob(rw http.ResponseWriter, r *http.Request) {
	var job models.Job

	vars := mux.Vars(r)

	jobID, err := strconv.Atoi(vars["JobID"])
	if err != nil {
		helpers.ErrorResponse(rw, errors.New("Invalid job ID"))
		return
	}

	findJob := db.Instance.Preload("Unit").Preload("Pot").First(&job, jobID)
	if findJob.Error != nil {
		helpers.ErrorResponse(rw, errors.New("Non-existing job ID"))
		return
	}

	helpers.SuccessResponse(rw, job)
}

func getJobs(rw http.ResponseWriter, r *http.Request) {
	var jobs []models.Job

	getJobs := db.Instance.Preload("Unit").Preload("Pot").Find(&jobs)

	if getJobs.Error != nil {
		helpers.ErrorResponse(rw, getJobs.Error)
		return
	}

	helpers.SuccessResponse(rw, jobs)
}

func updateJob(rw http.ResponseWriter, r *http.Request) {
	var job models.Job

	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data", err)
		return
	}

	jobID := r.PostForm.Get("JobID")
	startTimeStr := r.PostForm.Get("StartTime")
	statusStr := r.PostForm.Get("Status")
	waterQtyStr := r.PostForm.Get("WaterQty")

	if !helpers.CheckParams(jobID, waterQtyStr, startTimeStr, statusStr) {
		helpers.ErrorResponse(rw, errors.New("Invalid request. Missing parameters"))
		return
	}

	startTimeInt, _ := strconv.ParseInt(startTimeStr, 10, 64)
	startTime := time.Unix(startTimeInt, 0)
	waterQty, _ := strconv.Atoi(waterQtyStr)
	status, _ := strconv.Atoi(statusStr)

	if startTime.Before(time.Now()) {
		helpers.ErrorResponse(rw, errors.New("Start time should be in the future"))
		return
	}

	findJob := db.Instance.First(&job, jobID)
	if findJob.Error != nil {
		helpers.ErrorResponse(rw, findJob.Error)
		return
	}

	db.Instance.Model(&job).Updates(models.Job{
		WaterQty:  waterQty,
		StartTime: startTime,
		Status:    status,
	})

	helpers.SuccessResponse(rw, &job)
}

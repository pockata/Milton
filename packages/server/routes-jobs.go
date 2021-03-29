package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

	if !checkParams(unitID, potID, waterQtyStr, startTimeStr, statusStr) {
		errorResponse(rw, errors.New("Invalid request. Missing parameters"))
		return
	}

	startTimeInt, _ := strconv.ParseInt(startTimeStr, 10, 64)
	startTime := time.Unix(startTimeInt, 0)
	waterQty, _ := strconv.Atoi(waterQtyStr)
	status, _ := strconv.Atoi(statusStr)

	if startTime.Before(time.Now()) {
		errorResponse(rw, errors.New("Start time should be in the future"))
		return
	}

	var unit Unit
	var pot Pot

	findUnit := db.instance.First(&unit, unitID)
	if findUnit.Error != nil {
		errorResponse(rw, findUnit.Error)
		return
	}

	findPot := db.instance.First(&pot, potID)
	if findPot.Error != nil {
		errorResponse(rw, findPot.Error)
		return
	}

	entry := &Job{
		Unit:      unit,
		Pot:       pot,
		WaterQty:  waterQty,
		StartTime: startTime,
		Status:    status,
	}

	createEntry(rw, r, *db.instance, &entry)
}

func removeJob(rw http.ResponseWriter, r *http.Request) {
	deleteEntry(rw, r, *db.instance, &Job{})
}

func getJob(rw http.ResponseWriter, r *http.Request) {
	var job Job

	vars := mux.Vars(r)

	jobID, err := strconv.Atoi(vars["JobID"])
	if err != nil {
		errorResponse(rw, errors.New("Invalid job ID"))
		return
	}

	findJob := db.instance.Preload("Unit").Preload("Pot").First(&job, jobID)
	if findJob.Error != nil {
		errorResponse(rw, errors.New("Non-existing job ID"))
		return
	}

	successResponse(rw, job)
}

func getJobs(rw http.ResponseWriter, r *http.Request) {
	var jobs []Job

	getJobs := db.instance.Preload("Unit").Preload("Pot").Find(&jobs)

	if getJobs.Error != nil {
		errorResponse(rw, getJobs.Error)
		return
	}

	successResponse(rw, jobs)
}

func updateJob(rw http.ResponseWriter, r *http.Request) {
	var job Job

	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data", err)
		return
	}

	jobID := r.PostForm.Get("JobID")
	startTimeStr := r.PostForm.Get("StartTime")
	statusStr := r.PostForm.Get("Status")
	waterQtyStr := r.PostForm.Get("WaterQty")

	if !checkParams(jobID, waterQtyStr, startTimeStr, statusStr) {
		errorResponse(rw, errors.New("Invalid request. Missing parameters"))
		return
	}

	startTimeInt, _ := strconv.ParseInt(startTimeStr, 10, 64)
	startTime := time.Unix(startTimeInt, 0)
	waterQty, _ := strconv.Atoi(waterQtyStr)
	status, _ := strconv.Atoi(statusStr)

	if startTime.Before(time.Now()) {
		errorResponse(rw, errors.New("Start time should be in the future"))
		return
	}

	findJob := db.instance.First(&job, jobID)
	if findJob.Error != nil {
		errorResponse(rw, findJob.Error)
		return
	}

	db.instance.Model(&job).Updates(Job{
		WaterQty:  waterQty,
		StartTime: startTime,
		Status:    status,
	})

	successResponse(rw, &job)
}

package app

import (
	"fmt"
	"milton"
)

func (a App) AddJob(cfg milton.JobCreateConfig) (milton.Job, error) {
	job, err := a.jobService.Add(cfg)

	if err != nil {
		return nil, fmt.Errorf("couldn't add job: %w", err)
	}

	return job, nil
}

func (a App) RemoveJob(ID string) error {
	err := a.jobService.Remove(ID)

	if err != nil {
		return fmt.Errorf("couldn't remove job: %w", err)
	}

	return nil
}

func (a App) UpdateJob(ID string, cfg milton.JobUpdateConfig) (milton.Job, error) {
	job, err := a.jobService.Update(ID, cfg)
	if err != nil {
		return nil, fmt.Errorf("couldn't update job: %w", err)
	}

	return job, err
}

func (a App) GetJob(ID string) (milton.Job, error) {
	job, err := a.jobService.Get(ID)
	if err != nil {
		return nil, fmt.Errorf("couldn't get job: %w", err)
	}

	return job, nil
}

func (a App) GetAllJobs() (milton.JobSlice, error) {
	jobs, err := a.jobService.GetAll()
	if err != nil {
		return nil, fmt.Errorf("couldn't get all jobs: %w", err)
	}

	return jobs, nil
}

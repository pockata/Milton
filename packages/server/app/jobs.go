package app

import (
	"fmt"
	"milton"
)

func (a App) AddJob(cfg milton.JobCreateConfig) (milton.Job, error) {
	job, err := a.jobRepository.Add(cfg)

	if err != nil {
		return nil, fmt.Errorf("couldn't add job: %w", err)
	}

	return job, nil
}

func (a App) RemoveJob(ID string) error {
	err := a.jobRepository.Remove(ID)

	if err != nil {
		return fmt.Errorf("couldn't remove job: %w", err)
	}

	return nil
}

func (a App) UpdateJob(ID string, cfg milton.JobUpdateConfig) (milton.Job, error) {
	job, err := a.jobRepository.Update(ID, cfg)
	if err != nil {
		return nil, fmt.Errorf("couldn't update job: %w", err)
	}

	return job, err
}

func (a App) GetJob(ID string) (milton.Job, error) {
	job, err := a.jobRepository.Get(ID)
	if err != nil {
		return nil, fmt.Errorf("couldn't get job: %w", err)
	}

	return job, nil
}

func (a App) GetAllJobs() (milton.JobSlice, error) {
	jobs, err := a.jobRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("couldn't get all jobs: %w", err)
	}

	return jobs, nil
}

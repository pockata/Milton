package services

import (
	"fmt"
	"milton/core/domain"
	"milton/core/ports"
)

type JobService struct {
	jobs ports.JobRepository
}

func NewJobService(jobs ports.JobRepository) JobService {
	return JobService{
		jobs: jobs,
	}
}

func (s JobService) AddJob(cfg ports.JobCreateConfig) (domain.Job, error) {
	job, err := s.jobs.Add(cfg)

	if err != nil {
		return nil, fmt.Errorf("couldn't add job: %w", err)
	}

	return job, nil
}

func (s JobService) RemoveJob(ID string) error {
	err := s.jobs.Remove(ID)

	if err != nil {
		return fmt.Errorf("couldn't remove job: %w", err)
	}

	return nil
}

func (s JobService) UpdateJob(ID string, cfg ports.JobUpdateConfig) (domain.Job, error) {
	job, err := s.jobs.Update(ID, cfg)
	if err != nil {
		return nil, fmt.Errorf("couldn't update job: %w", err)
	}

	return job, err
}

func (s JobService) GetJob(ID string) (domain.Job, error) {
	job, err := s.jobs.Get(ID)
	if err != nil {
		return nil, fmt.Errorf("couldn't get job: %w", err)
	}

	return job, nil
}

func (s JobService) GetAllJobs() (domain.JobSlice, error) {
	jobs, err := s.jobs.GetAll()
	if err != nil {
		return nil, fmt.Errorf("couldn't get all jobs: %w", err)
	}

	return jobs, nil
}

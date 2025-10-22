package service

import (
	"context"
	"next-go-sample/internal/domain"
	"next-go-sample/internal/repository"
)

// SampleService handles business logic for samples
type SampleService struct {
	repo repository.SampleRepository
}

// NewSampleService creates a new SampleService
func NewSampleService(repo repository.SampleRepository) *SampleService {
	return &SampleService{
		repo: repo,
	}
}

// GetAllSamples returns all samples
func (s *SampleService) GetAllSamples(ctx context.Context) ([]*domain.Sample, error) {
	return s.repo.FindAll(ctx)
}

// GetSampleByID returns a sample by ID
func (s *SampleService) GetSampleByID(ctx context.Context, id string) (*domain.Sample, error) {
	return s.repo.FindByID(ctx, id)
}

// CreateSample creates a new sample
func (s *SampleService) CreateSample(ctx context.Context, name string) (*domain.Sample, error) {
	sample := domain.NewSample(name)

	if err := s.repo.Create(ctx, sample); err != nil {
		return nil, err
	}

	return sample, nil
}

// UpdateSample updates an existing sample
func (s *SampleService) UpdateSample(ctx context.Context, id string, name string) (*domain.Sample, error) {
	sample, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	sample.UpdateName(name)

	if err := s.repo.Update(ctx, sample); err != nil {
		return nil, err
	}

	return sample, nil
}

// DeleteSample deletes a sample by ID
func (s *SampleService) DeleteSample(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

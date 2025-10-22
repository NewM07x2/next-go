package repository

import (
	"context"
	"next-go-task/internal/domain"
)

// SampleRepository defines the interface for sample data access
type SampleRepository interface {
	// FindAll returns all samples
	FindAll(ctx context.Context) ([]*domain.Sample, error)

	// FindByID returns a sample by ID
	FindByID(ctx context.Context, id string) (*domain.Sample, error)

	// Create creates a new sample
	Create(ctx context.Context, sample *domain.Sample) error

	// Update updates an existing sample
	Update(ctx context.Context, sample *domain.Sample) error

	// Delete deletes a sample by ID
	Delete(ctx context.Context, id string) error
}

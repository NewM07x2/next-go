package repository

import (
	"context"
	"fmt"
	"next-go-task/internal/domain"
	"strconv"
	"sync"
	"time"
)

// InMemorySampleRepository is an in-memory implementation of SampleRepository
// TODO: Replace with actual database implementation
type InMemorySampleRepository struct {
	mu      sync.RWMutex
	samples map[string]*domain.Sample
	nextID  int
}

// NewInMemorySampleRepository creates a new in-memory sample repository
func NewInMemorySampleRepository() *InMemorySampleRepository {
	repo := &InMemorySampleRepository{
		samples: make(map[string]*domain.Sample),
		nextID:  1,
	}

	// Initialize with sample data
	repo.initializeSampleData()

	return repo
}

func (r *InMemorySampleRepository) initializeSampleData() {
	sample1 := &domain.Sample{
		ID:        "1",
		Name:      "Sample 1",
		CreatedAt: time.Now().Add(-24 * time.Hour),
		UpdatedAt: time.Now().Add(-24 * time.Hour),
	}

	sample2 := &domain.Sample{
		ID:        "2",
		Name:      "Sample 2",
		CreatedAt: time.Now().Add(-12 * time.Hour),
		UpdatedAt: time.Now().Add(-12 * time.Hour),
	}

	r.samples["1"] = sample1
	r.samples["2"] = sample2
	r.nextID = 3
}

// FindAll returns all samples
func (r *InMemorySampleRepository) FindAll(ctx context.Context) ([]*domain.Sample, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	samples := make([]*domain.Sample, 0, len(r.samples))
	for _, sample := range r.samples {
		samples = append(samples, sample)
	}

	return samples, nil
}

// FindByID returns a sample by ID
func (r *InMemorySampleRepository) FindByID(ctx context.Context, id string) (*domain.Sample, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	sample, exists := r.samples[id]
	if !exists {
		return nil, fmt.Errorf("sample not found: %s", id)
	}

	return sample, nil
}

// Create creates a new sample
func (r *InMemorySampleRepository) Create(ctx context.Context, sample *domain.Sample) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	sample.ID = strconv.Itoa(r.nextID)
	r.nextID++

	r.samples[sample.ID] = sample

	return nil
}

// Update updates an existing sample
func (r *InMemorySampleRepository) Update(ctx context.Context, sample *domain.Sample) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.samples[sample.ID]; !exists {
		return fmt.Errorf("sample not found: %s", sample.ID)
	}

	r.samples[sample.ID] = sample

	return nil
}

// Delete deletes a sample by ID
func (r *InMemorySampleRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.samples[id]; !exists {
		return fmt.Errorf("sample not found: %s", id)
	}

	delete(r.samples, id)

	return nil
}

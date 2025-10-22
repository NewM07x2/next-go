package domain

import "time"

// Sample represents the core business entity for a sample
type Sample struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewSample creates a new Sample with default values
func NewSample(name string) *Sample {
	now := time.Now()
	return &Sample{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// UpdateName updates the sample name
func (s *Sample) UpdateName(name string) {
	s.Name = name
	s.UpdatedAt = time.Now()
}

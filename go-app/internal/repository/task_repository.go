package repository

import (
	"context"
	"next-go-task/internal/domain"
)

// TaskRepository defines the interface for task data access
type TaskRepository interface {
	// FindAll returns all tasks
	FindAll(ctx context.Context) ([]*domain.Task, error)

	// FindByID returns a task by ID
	FindByID(ctx context.Context, id string) (*domain.Task, error)

	// Create creates a new task
	Create(ctx context.Context, task *domain.Task) error

	// Update updates an existing task
	Update(ctx context.Context, task *domain.Task) error

	// Delete deletes a task by ID
	Delete(ctx context.Context, id string) error
}

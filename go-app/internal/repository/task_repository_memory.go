package repository

import (
	"context"
	"fmt"
	"next-go-task/internal/domain"
	"strconv"
	"sync"
	"time"
)

// InMemoryTaskRepository is an in-memory implementation of TaskRepository
// TODO: Replace with actual database implementation
type InMemoryTaskRepository struct {
	mu     sync.RWMutex
	tasks  map[string]*domain.Task
	nextID int
}

// NewInMemoryTaskRepository creates a new in-memory task repository
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	repo := &InMemoryTaskRepository{
		tasks:  make(map[string]*domain.Task),
		nextID: 1,
	}

	// Initialize with sample data
	repo.initializeSampleData()

	return repo
}

func (r *InMemoryTaskRepository) initializeSampleData() {
	desc1 := "Need to buy milk from the store"
	desc2 := "Write the monthly report for the team"

	task1 := &domain.Task{
		ID:          "1",
		Title:       "Buy milk",
		Description: &desc1,
		Completed:   false,
		CreatedAt:   time.Now().Add(-24 * time.Hour),
		UpdatedAt:   time.Now().Add(-24 * time.Hour),
	}

	task2 := &domain.Task{
		ID:          "2",
		Title:       "Write report",
		Description: &desc2,
		Completed:   false,
		CreatedAt:   time.Now().Add(-12 * time.Hour),
		UpdatedAt:   time.Now().Add(-12 * time.Hour),
	}

	r.tasks["1"] = task1
	r.tasks["2"] = task2
	r.nextID = 3
}

// FindAll returns all tasks
func (r *InMemoryTaskRepository) FindAll(ctx context.Context) ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*domain.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// FindByID returns a task by ID
func (r *InMemoryTaskRepository) FindByID(ctx context.Context, id string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, fmt.Errorf("task not found: %s", id)
	}

	return task, nil
}

// Create creates a new task
func (r *InMemoryTaskRepository) Create(ctx context.Context, task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = strconv.Itoa(r.nextID)
	r.nextID++

	r.tasks[task.ID] = task

	return nil
}

// Update updates an existing task
func (r *InMemoryTaskRepository) Update(ctx context.Context, task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return fmt.Errorf("task not found: %s", task.ID)
	}

	r.tasks[task.ID] = task

	return nil
}

// Delete deletes a task by ID
func (r *InMemoryTaskRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return fmt.Errorf("task not found: %s", id)
	}

	delete(r.tasks, id)

	return nil
}

package service

import (
	"context"
	"next-go-task/internal/domain"
	"next-go-task/internal/repository"
)

// TaskService handles business logic for tasks
type TaskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a new TaskService
func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// GetAllTasks returns all tasks
func (s *TaskService) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	return s.repo.FindAll(ctx)
}

// GetTaskByID returns a task by ID
func (s *TaskService) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	return s.repo.FindByID(ctx, id)
}

// CreateTask creates a new task
func (s *TaskService) CreateTask(ctx context.Context, title string, description *string) (*domain.Task, error) {
	task := domain.NewTask(title, description)

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

// UpdateTask updates an existing task
func (s *TaskService) UpdateTask(ctx context.Context, id string, title *string, description *string, completed *bool) (*domain.Task, error) {
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	task.Update(title, description, completed)

	if err := s.repo.Update(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

// DeleteTask deletes a task by ID
func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

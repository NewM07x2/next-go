package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"next-go-task/internal/service"
)

type Resolver struct {
	TaskService *service.TaskService
}

func NewResolver(taskService *service.TaskService) *Resolver {
	return &Resolver{
		TaskService: taskService,
	}
}

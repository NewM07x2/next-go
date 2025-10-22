package model

// Task represents a task item
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// CreateTaskRequest represents the request body for creating a task
type CreateTaskRequest struct {
	Title string `json:"title" validate:"required"`
}

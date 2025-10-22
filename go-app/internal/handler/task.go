package handler

import (
	"net/http"

	"next-go-task/internal/model"

	"github.com/labstack/echo/v4"
)

// TaskHandler handles task-related requests
type TaskHandler struct{}

// NewTaskHandler creates a new TaskHandler
func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

// GetAll returns all tasks
func (h *TaskHandler) GetAll(c echo.Context) error {
	// TODO: Replace with actual database query
	tasks := []model.Task{
		{ID: 1, Title: "Buy milk"},
		{ID: 2, Title: "Write report"},
	}
	return c.JSON(http.StatusOK, tasks)
}

// Create creates a new task
func (h *TaskHandler) Create(c echo.Context) error {
	var req model.CreateTaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Title is required",
		})
	}

	// TODO: Replace with actual database insert
	created := model.Task{
		ID:    3,
		Title: req.Title,
	}

	return c.JSON(http.StatusCreated, created)
}

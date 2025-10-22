package domain

import "time"

// Task represents the core business entity for a task
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// NewTask creates a new Task with default values
func NewTask(title string, description *string) *Task {
	now := time.Now()
	return &Task{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// MarkAsCompleted marks the task as completed
func (t *Task) MarkAsCompleted() {
	t.Completed = true
	t.UpdatedAt = time.Now()
}

// Update updates the task fields
func (t *Task) Update(title *string, description *string, completed *bool) {
	if title != nil {
		t.Title = *title
	}
	if description != nil {
		t.Description = description
	}
	if completed != nil {
		t.Completed = *completed
	}
	t.UpdatedAt = time.Now()
}

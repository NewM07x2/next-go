package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"next-go-task/internal/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTasks(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewTaskHandler()

	// Test
	if assert.NoError(t, handler.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var tasks []model.Task
		err := json.Unmarshal(rec.Body.Bytes(), &tasks)
		assert.NoError(t, err)
		assert.Len(t, tasks, 2)
		assert.Equal(t, "Buy milk", tasks[0].Title)
	}
}

func TestCreateTask(t *testing.T) {
	// Setup
	e := echo.New()
	payload := `{"title":"New Task"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewTaskHandler()

	// Test
	if assert.NoError(t, handler.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var task model.Task
		err := json.Unmarshal(rec.Body.Bytes(), &task)
		assert.NoError(t, err)
		assert.Equal(t, "New Task", task.Title)
	}
}

func TestCreateTask_InvalidPayload(t *testing.T) {
	// Setup
	e := echo.New()
	payload := `{"title":""}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewTaskHandler()

	// Test
	if assert.NoError(t, handler.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Title is required", response["error"])
	}
}

func TestCreateTask_MalformedJSON(t *testing.T) {
	// Setup
	e := echo.New()
	payload := `{invalid json}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewTaskHandler()

	// Test
	if assert.NoError(t, handler.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

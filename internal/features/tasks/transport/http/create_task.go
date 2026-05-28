package tasks_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/0ScPro0/go-todolist/internal/core/domain"
	"github.com/0ScPro0/go-todolist/internal/core/logger"
	"go.uber.org/zap"
)

type CreateTaskRequest struct {
	Name        string            `json:"name"`
	Description *string    		  `json:"description"`
	Status      domain.TaskStatus `json:"status"`
}

type CreateTaskResponse struct {
	ID      int `json:"id"`
	Version int `json:"version"`

	Name        string     		  `json:"name"`
	Description *string    		  `json:"description"`
	Status      domain.TaskStatus `json:"status"`
}

func (h *TasksHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)

	log.Debug("Invoke CreateTask handler")
	var request CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error("CreateTask:", zap.Error(err))
	}

	rw.WriteHeader(http.StatusOK)
}
package tasks_transport_http

import (
	"net/http"

	domain "github.com/0ScPro0/go-todolist/internal/core/domain"
	core_http_server "github.com/0ScPro0/go-todolist/internal/core/transport/http/server"
)

type TasksHTTPHandler struct {
	taskService TasksService
}

type TasksService interface {
	Create(t domain.Task)
}

func NewTaskHTTPHandler(
	taskService TasksService,
) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		taskService: taskService,
	}
}

func (t *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodPost,
			Path: "/tasks",
			Handler: t.CreateTask,
		},
	}
}
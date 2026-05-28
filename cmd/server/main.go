package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/0ScPro0/go-todolist/internal/core/config"
	"github.com/0ScPro0/go-todolist/internal/core/logger"
	core_http_middleware "github.com/0ScPro0/go-todolist/internal/core/transport/http/middleware"
	core_http_server "github.com/0ScPro0/go-todolist/internal/core/transport/http/server"
	tasks_transport_http "github.com/0ScPro0/go-todolist/internal/features/tasks/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	log, err := logger.NewLogger(cfg)
	if err != nil {
		panic(err)
	}
	defer log.Close()
	log.Info("Logger successfully initialized")

	tasksTransportHTTP := tasks_transport_http.NewTaskHTTPHandler(nil)
	tasksRoutes := tasksTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVerssionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(tasksRoutes...)

	httpServer := core_http_server.NewHTTPServer(
		cfg, 
		log,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(log),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	httpServer.RegisterAPIVersionRouter(*apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server run error:", zap.Error(err))
	}
}

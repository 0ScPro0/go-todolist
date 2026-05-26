package main

import (
	"github.com/0ScPro0/go-todolist/internal/core/config"
	"github.com/0ScPro0/go-todolist/internal/core/logger"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	log, err := logger.NewLogger(cfg)
	if err != nil {
		panic(err)
	}
	log.Info("Logger successfully initialized")
}

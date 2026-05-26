package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	Host string `envconfig:"SERVER_HOST" required:"true"`
	Port int `envconfig:"SERVER_PORT" required:"true"`
}

type DatabaseConfig struct {
	DbUrl string `envconfig:"DATABASE_POSTGRES_URL" required:"true"`
}

type LoggerConfig struct {
	Level string `envconfig:"LOGGER_LEVEL" required:"true"`
	Folder string `envconfig:"LOGGER_FOLDER" required:"true"`
}

type EnvironmentConfig struct {
	Debug bool `envconfig:"ENVIRONMENT_DEBUG" required:"false"`
}

type Config struct {
	Server      ServerConfig
	Database    DatabaseConfig
	Logger 		LoggerConfig
	Environment EnvironmentConfig
}

func NewConfig() (*Config, error) {
	var config Config

	processors := map[string]interface{}{
		"SERVER":      &config.Server,
		"DATABASE":    &config.Database,
		"LOGGER":      &config.Logger,
		"ENVIRONMENT": &config.Environment,
	}
	
	for prefix, cfg := range processors {
		if err := envconfig.Process(prefix, cfg); err != nil {
			return nil, fmt.Errorf("unable to get %s config: %w", prefix, err)
		}
	}
	
	return &config, nil
}
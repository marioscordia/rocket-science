package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type loggerEnvConfig struct {
	Level  string `env:"LOGGER_LEVEL,required"`
	AsJson bool   `env:"LOGGER_AS_JSON,required"`
}

type loggerConfig struct {
	raw loggerEnvConfig
}

func NewLoggerConfig() (*loggerConfig, error) {
	var raw loggerEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, fmt.Errorf("error parsing for Logger config: %w", err)
	}

	return &loggerConfig{raw: raw}, nil
}

func (cfg *loggerConfig) GetLevel() string {
	return cfg.raw.Level
}

func (cfg *loggerConfig) AsJSON() bool {
	return cfg.raw.AsJson
}

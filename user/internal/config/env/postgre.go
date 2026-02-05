package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type postgreEnvConfig struct {
	Host     string `env:"POSTGRES_HOST,required"`
	Port     int    `env:"POSTGRES_PORT,required"`
	User     string `env:"POSTGRES_USER,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	Database string `env:"POSTGRES_DB,required"`
	SSLMode  string `env:"POSTGRES_SSL_MODE,required"`
}

type postgreConfig struct {
	raw postgreEnvConfig
}

func NewPostgreConfig() (*postgreConfig, error) {
	var raw postgreEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, fmt.Errorf("error parsing for Postgre config")
	}

	return &postgreConfig{raw: raw}, nil
}

func (cfg *postgreConfig) GetURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.raw.User,
		cfg.raw.Password,
		cfg.raw.Host,
		cfg.raw.Port,
		cfg.raw.Database,
		cfg.raw.SSLMode,
	)
}

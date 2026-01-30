package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type httpEnvConfig struct {
	Port        int    `env:"HTTP_PORT,required"`
	Host        string `env:"HTTP_HOST" envDefault:"localhost"`
	ReadTimeout int    `env:"HTTP_READ_TIMEOUT" envDefault:"15"`
}

type httpConfig struct {
	raw httpEnvConfig
}

func NewHTTPConfig() (*httpConfig, error) {
	var raw httpEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, fmt.Errorf("error parsing for HTTP config: %w", err)
	}

	return &httpConfig{raw: raw}, nil
}

func (c *httpConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", c.raw.Host, c.raw.Port)
}

func (c *httpConfig) GetReadTimeout() int {
	return c.raw.ReadTimeout
}

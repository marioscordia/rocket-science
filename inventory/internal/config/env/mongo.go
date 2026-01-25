package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type mongoEnvConfig struct {
	Host       string `env:"MONGO_HOST,required"`
	Port       int    `env:"MONGO_PORT,required"`
	User       string `env:"MONGO_USER,required"`
	Password   string `env:"MONGO_PASSWORD,required"`
	Database   string `env:"MONGO_DATABASE,required"`
	Collection string `env:"MONGO_COLLECTION,required"`
}

type mongoConfig struct {
	raw mongoEnvConfig
}

func NewMongoConfig() (*mongoConfig, error) {
	var raw mongoEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, fmt.Errorf("error parsing for Mongo config: %w", err)
	}

	return &mongoConfig{raw: raw}, nil
}

func (cfg *mongoConfig) GetURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d",
		cfg.raw.User,
		cfg.raw.Password,
		cfg.raw.Host,
		cfg.raw.Port,
	)
}

func (cfg *mongoConfig) GetDatabase() string {
	return cfg.raw.Database
}
func (cfg *mongoConfig) GetCollection() string {
	return cfg.raw.Collection
}

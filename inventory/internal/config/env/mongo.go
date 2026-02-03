package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type mongoEnvConfig struct {
	Host               string `env:"MONGO_HOST,required"`
	Port               int    `env:"MONGO_PORT,required"`
	InitDBRootUsername string `env:"MONGO_INITDB_ROOT_USERNAME,required"`
	InitDBRootPassword string `env:"MONGO_INITDB_ROOT_PASSWORD,required"`
	Database           string `env:"MONGO_DATABASE,required"`
	Collection         string `env:"MONGO_COLLECTION,required"`
	AuthDB             string `env:"MONGO_AUTH_DB,required"`
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
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		cfg.raw.InitDBRootUsername,
		cfg.raw.InitDBRootPassword,
		cfg.raw.Host,
		cfg.raw.Port,
		cfg.raw.AuthDB,
	)
}

func (cfg *mongoConfig) GetDatabase() string {
	return cfg.raw.Database
}
func (cfg *mongoConfig) GetCollection() string {
	return cfg.raw.Collection
}

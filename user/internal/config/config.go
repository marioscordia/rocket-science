package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/marioscordia/rocket-science/user/internal/config/env"
)

var appConfig *config

type config struct {
	Logger   LoggerConfig
	GRPC     GRPCConfig
	Redis    RedisConfig
	Postgres PostgresConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	grpcCfg, err := env.NewGRPCConfig()
	if err != nil {
		return err
	}

	postgresCfg, err := env.NewPostgreConfig()
	if err != nil {
		return err
	}

	redisCfg, err := env.NewRedisConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:   loggerCfg,
		GRPC:     grpcCfg,
		Postgres: postgresCfg,
		Redis:    redisCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}

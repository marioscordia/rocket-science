package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/marioscordia/rocket-science/order/internal/config/env"
)

var appConfig *config

type config struct {
	Logger  LoggerConfig
	GRPC    GRPCConfig
	Postgre PostgresConfig
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

	grpcCfg, err := env.NewOrderGRPCConfig()
	if err != nil {
		return err
	}

	postgreCfg, err := env.NewPostgreConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:  loggerCfg,
		GRPC:    grpcCfg,
		Postgre: postgreCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}

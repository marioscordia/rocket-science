package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/marioscordia/rocket-science/inventory/internal/config/env"
)

var appConfig *config

type config struct {
	Logger       LoggerConfig
	GRPC         GRPCConfig
	Mongo        MongoConfig
	InventorySvc GRPCConfig
	PaymentSvc   GRPCConfig
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

	inventorySvcCfg, err := env.NewInventorySvcConfig()
	if err != nil {
		return err
	}

	paymentSvcCfg, err := env.NewPaymentSvcConfig()
	if err != nil {
		return err
	}

	mongoCfg, err := env.NewMongoConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:       loggerCfg,
		GRPC:         grpcCfg,
		Mongo:        mongoCfg,
		InventorySvc: inventorySvcCfg,
		PaymentSvc:   paymentSvcCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}

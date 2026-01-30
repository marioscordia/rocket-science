package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/marioscordia/rocket-science/order/internal/config/env"
)

var appConfig *config

type config struct {
	Logger       LoggerConfig
	GRPC         GRPCConfig
	HTTP         HTTPConfig
	Postgre      PostgresConfig
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

	httpCfg, err := env.NewHTTPConfig()
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

	postgreCfg, err := env.NewPostgreConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:       loggerCfg,
		HTTP:         httpCfg,
		Postgre:      postgreCfg,
		InventorySvc: inventorySvcCfg,
		PaymentSvc:   paymentSvcCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}

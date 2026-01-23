package env

import (
	"fmt"
	"net"

	"github.com/caarlos0/env/v11"
)

type paymentSvcConfig struct {
	Host string `env:"PAYMENT_GRPC_HOST,required"`
	Port string `env:"PAYMENT_GRPC_PORT,required"`
}

func NewPaymentSvcConfig() (*paymentSvcConfig, error) {
	var raw paymentSvcConfig
	if err := env.Parse(&raw); err != nil {
		return nil, fmt.Errorf("error parsing for Payment Service config: %w", err)
	}

	return &raw, nil
}

func (cfg *paymentSvcConfig) GetAddress() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

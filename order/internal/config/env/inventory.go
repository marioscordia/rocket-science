package env

import (
	"fmt"
	"net"

	"github.com/caarlos0/env/v11"
)

type inventorySvcConfig struct {
	Host string `env:"INVENTORY_GRPC_HOST,required"`
	Port string `env:"INVENTORY_GRPC_PORT,required"`
}

func NewInventorySvcConfig() (*inventorySvcConfig, error) {
	var raw inventorySvcConfig
	if err := env.Parse(&raw); err != nil {
		return nil, fmt.Errorf("error parsing for Inventory Service config: %w", err)
	}

	return &raw, nil
}

func (cfg *inventorySvcConfig) GetAddress() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

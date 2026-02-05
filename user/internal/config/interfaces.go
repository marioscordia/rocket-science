package config

import "time"

type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
	CacheTTL() time.Duration
}

type GRPCConfig interface {
	Address() string
}

type LoggerConfig interface {
	Level() string
	AsJSON() bool
}

type PostgresConfig interface {
	GetURL() string
}

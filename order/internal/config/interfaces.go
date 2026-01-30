package config

type LoggerConfig interface {
	GetLevel() string
	AsJSON() bool
}

type GRPCConfig interface {
	GetAddress() string
}

type PostgresConfig interface {
	GetURL() string
}

type HTTPConfig interface {
	GetAddress() string
	GetReadTimeout() int
}

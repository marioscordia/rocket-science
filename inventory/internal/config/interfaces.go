package config

type LoggerConfig interface {
	GetLevel() string
	AsJSON() bool
}

type GRPCConfig interface {
	GetAddress() string
}

type MongoConfig interface {
	GetURI() string
	GetDatabase() string
	GetCollection() string
}

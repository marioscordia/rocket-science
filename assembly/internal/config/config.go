package config

var appConfig *config

type config struct {
}

func AppConfig() *config {
	return appConfig
}

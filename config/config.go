package config

import (
	"github.com/caarlos0/env"
)

// Conf holds the environment configurations taken from docker-compose.yml
type Conf struct {
	SampleEnvVariable     string `env:"SAMPLE_ENV_VARIABLE"`
}

// Config is a global configuration
var Config Conf

// InitEnv grabs the environment variables found within the docker-compose.yml file
func InitEnv() Conf {
	if err := env.Parse(&Config); err != nil {
		panic(err)
	}
	return Config
}
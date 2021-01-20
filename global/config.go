package global

import "github.com/kelseyhightower/envconfig"

type Config struct {
	AppName string `envconfig:"APP_NAME" default:"user-api"`
	Env     string `envconfig:"ENV" default:"dev"`
	Port    string    `envconfig:"PORT" default:"8080""`
}

func New() (*Config, error) {
	config := new(Config)
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}

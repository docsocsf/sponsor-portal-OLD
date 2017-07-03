package config

import (
	"github.com/caarlos0/env"
)

type HostConfig struct {
	Port        string `env:"PORT" envDefault:"8080"`
	StaticFiles string `env:"STATIC_FILES,required"`
}

func GetHost() (HostConfig, error) {
	var host HostConfig
	err := env.Parse(&host)
	host.Port = ":" + host.Port
	return host, err
}

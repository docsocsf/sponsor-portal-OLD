package config

import "github.com/caarlos0/env"

type Auth struct {
	BaseURL string `env:"AUTH_BASE_URL"`
}

type BasicAuth struct {
	ServiceUsername string `env:"SERVICE_USER_NAME,required"`
	ServicePassword string `env:"SERVICE_PASSWORD,required"`
	Realm string `env:"REALM,required"`
}

func GetAuth() (auth Auth, err error) {
	err = env.Parse(&auth)
	return
}

func GetBasicAuth() (auth BasicAuth, err error) {
	err = env.Parse(&auth)
	return
}

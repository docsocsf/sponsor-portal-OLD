package config

import "github.com/caarlos0/env"

type Auth struct {
	CookieSecret string `env:"AUTH_COOKIE_SECRET,required"`

	BaseURL string `env:"AUTH_BASE_URL"`

	JwtSecret string `env:"JWT_SECRET,required"`
	JwtIssuer string `env:"JWT_ISSUER,required"`
}

type BasicAuth struct {
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

package config

import "github.com/caarlos0/env"

type OAuth struct {
	CookieSecret string `env:"OAUTH_COOKIE_SECRET"`

	BaseURL string `env:"OAUTH_BASE_URL"`
	Issuer  string `env:"OAUTH_ISSUER"`

	ClientID     string `env:"OAUTH_CLIENT_ID"`
	ClientSecret string `env:"OAUTH_CLIENT_SECRET"`
}

func GetAuth() (auth OAuth, err error) {
	err = env.Parse(&auth)
	return
}

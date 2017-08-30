package config

import "github.com/caarlos0/env"

type OAuth struct {
	CookieSecret string `env:"OAUTH_COOKIE_SECRET,required"`

	BaseURL string `env:"OAUTH_BASE_URL"`
	Issuer  string `env:"OAUTH_ISSUER"`

	ClientID     string `env:"OAUTH_CLIENT_ID,required"`
	ClientSecret string `env:"OAUTH_CLIENT_SECRET,required"`

	JwtSecret string `env:"JWT_SECRET,required"`
	JwtIssuer string `env:"JWT_ISSUER,required"`
}

func GetAuth() (auth OAuth, err error) {
	err = env.Parse(&auth)
	return
}

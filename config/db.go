package config

import "github.com/caarlos0/env"

type Database struct {
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASS"`

	DatabaseName string `env:"DB_NAME"`
	Host         string `env:"DB_HOST"`
	Port         string `env:"DB_PORT"`

	SslMode string `env:"DB_SSLMODE"`
}

func GetDB() (db Database, err error) {
	err = env.Parse(&db)
	return
}

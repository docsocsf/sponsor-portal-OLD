package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/docsocsf/sponsor-portal/config"
)

func NewDB(config config.Database) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Username, config.Password,
		config.Host, config.Port,
		config.DatabaseName,
		config.SslMode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

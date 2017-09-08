package sponsor

import (
	"net/http"
	"os"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/config"
	"github.com/docsocsf/sponsor-portal/model"
	"github.com/gorilla/handlers"
)

type Service struct {
	router http.Handler
	Auth   auth.Auth
	model.UserReader
}

func New(authConfig *auth.Config) (*Service, error) {
	service := Service{}

	if err := service.setupAuth(authConfig); err != nil {
		return nil, err
	}

	service.router = service.getRoutes()

	return &service, nil
}

func (s *Service) SetupDatabase(dbConfig config.Database) error {
	db, err := model.NewDB(dbConfig)
	if err != nil {
		return err
	}

	s.UserReader = model.NewUserReader(db)

	return nil
}

func (s *Service) Handler() http.Handler {
	return handlers.LoggingHandler(os.Stdout, s.router)
}

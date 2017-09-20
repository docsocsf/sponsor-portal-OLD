package sponsor

import (
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/config"
	"github.com/docsocsf/sponsor-portal/model"
	"github.com/gorilla/mux"
)

type Service struct {
	staticFiles string

	Auth auth.Auth
	model.UserReader
}

func New(authConfig *auth.Config, staticFiles string) (*Service, error) {
	service := Service{
		staticFiles: staticFiles,
	}

	if err := service.setupAuth(authConfig); err != nil {
		return nil, err
	}

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

func (s *Service) Handle(r *mux.Router, web http.Handler) {
	s.defineRoutes(r.PathPrefix("/sponsors").Subrouter(), web)
}

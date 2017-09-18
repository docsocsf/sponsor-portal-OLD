package student

import (
	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/config"
	"github.com/docsocsf/sponsor-portal/model"
	"github.com/gorilla/mux"
)

type Service struct {
	staticFiles string

	Auth auth.Auth
	s3   *model.S3

	model.UserReader
	model.CVReader
	model.CVWriter
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

func (s *Service) SetupStorer(s3Config config.S3) {
	s.s3 = model.NewS3(s3Config.Aws, s3Config.Bucket, s3Config.Prefix)
}

func (s *Service) SetupDatabase(dbConfig config.Database) error {
	db, err := model.NewDB(dbConfig)
	if err != nil {
		return err
	}

	s.UserReader = model.NewUserReader(db)
	s.CVReader = model.NewCVReader(db)
	s.CVWriter = model.NewCVWriter(db)

	return nil
}

func (s *Service) Handle(r *mux.Router) {
	s.defineRoutes(r.PathPrefix("/students").Subrouter())
}

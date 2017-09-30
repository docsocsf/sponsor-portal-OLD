package sponsor

import (
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/gorilla/mux"
)

func (s *Service) GetApiRoutes() http.Handler {
	api := mux.NewRouter()
	get := api.Methods(http.MethodGet).Subrouter()

	get.Handle("/cvs",
		auth.RequireJWT(http.HandlerFunc(s.getCVs), s.Auth, role),
	)

	get.Handle("/cv/{id:[0-9]+}/download",
		auth.RequireOnetimeJWT(http.HandlerFunc(s.downloadCV), s.Auth, role),
	)

	return api
}

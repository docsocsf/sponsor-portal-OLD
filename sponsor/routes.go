package sponsor

import (
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/gorilla/mux"
)

func (s *Service) defineRoutes(r *mux.Router, web http.Handler) {
	// auth
	r.PathPrefix("/auth").Handler(http.StripPrefix("/sponsors/auth", s.Auth.Handler()))

	r.Handle("/", auth.RequireAuth(web, "/login", role))
	r.PathPrefix("/api").Handler(http.StripPrefix("/sponsors/api", s.getApiRoutes()))
}

func (s *Service) getApiRoutes() http.Handler {
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

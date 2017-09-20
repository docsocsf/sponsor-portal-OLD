package student

import (
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/gorilla/mux"
)

func (s *Service) defineRoutes(r *mux.Router, web http.Handler) {
	// auth
	r.PathPrefix("/auth").Handler(http.StripPrefix("/students/auth", s.Auth.Handler()))

	r.Handle("/", auth.RequireAuth("/students/auth/login", web))
	r.PathPrefix("/api").Handler(http.StripPrefix("/students/api", s.getApiRoutes()))
}

func (s *Service) getApiRoutes() http.Handler {
	api := mux.NewRouter()

	api.HandleFunc("/user", s.getUserInformation)
	api.HandleFunc("/cv", s.uploadCV).Methods(http.MethodPost)
	api.HandleFunc("/cv", s.getCV).Methods(http.MethodGet)

	return auth.RequireJWT(s.Auth, api)
}

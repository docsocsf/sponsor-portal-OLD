package sponsor

import (
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/gorilla/mux"
)

func (s *Service) defineRoutes(r *mux.Router, web http.Handler) {
	// auth
	r.PathPrefix("/auth").Handler(http.StripPrefix("/sponsors/auth", s.Auth.Handler()))

	r.Handle("/", auth.RequireAuth("/login", web))
}

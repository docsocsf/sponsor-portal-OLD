package sponsor

import (
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/gorilla/mux"
)

func (s *Service) defineRoutes(r *mux.Router) {
	// auth
	r.PathPrefix("/auth").Handler(http.StripPrefix("/sponsors/auth", s.Auth.Handler()))

	r.Handle("/", auth.RequireAuth(s.Auth, indexHandler(s)))
}

func indexHandler(s *Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, s.staticFiles+"/sponsors.html")
	})
}

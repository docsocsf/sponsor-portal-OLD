package sponsor

import (
	"fmt"
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/gorilla/mux"
)

func (s *Service) getRoutes() http.Handler {
	router := mux.NewRouter()

	// auth
	router.PathPrefix("/auth/").Handler(http.StripPrefix("/auth", s.Auth.Handler()))

	router.Handle("/", auth.RequireAuth(s.Auth, indexHandler(s)))

	return router
}

func indexHandler(s *Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := auth.User(r)

		user, err := s.UserReader.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Hello, %s", user.Name)
	})
}

package student

import (
	"fmt"
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/gorilla/mux"
)

func (service *Service) getRoutes() http.Handler {
	router := mux.NewRouter()

	// auth
	router.PathPrefix("/auth/").Handler(http.StripPrefix("/auth", service.Auth.Handler()))

	router.Handle("/", service.Auth.RequireAuth(indexHandler(service)))

	return router
}

func indexHandler(s *Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := auth.User(r)

		user, err := s.UserReader.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Hello, %s", user.Name)
	})
}

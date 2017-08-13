package student

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) getRoutes() http.Handler {
	router := mux.NewRouter()

	// auth
	router.PathPrefix("/auth/").Handler(http.StripPrefix("/auth", s.Auth.Handler()))

	router.Handle("/", s.Auth.RequireAuth(indexHandler(s)))
	router.PathPrefix("/api/").Handler(http.StripPrefix("/api", s.getApiRoutes()))

	return router
}

func (s *Service) getApiRoutes() http.Handler {
	api := mux.NewRouter()
	api.HandleFunc("/user", s.getUserInformation)
	api.HandleFunc("/cv", s.uploadCV).Methods(http.MethodPost)
	api.HandleFunc("/cv", s.getCV).Methods(http.MethodGet)

	return s.Auth.RequireJWT(api)
}

func indexHandler(s *Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, s.staticFiles+"/students.html")
	})
}

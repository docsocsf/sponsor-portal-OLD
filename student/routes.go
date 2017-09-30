package student

import (
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/gorilla/mux"
)

func (s *Service) GetApiRoutes() http.Handler {
	api := mux.NewRouter()

	api.HandleFunc("/user", s.getUserInformation)
	api.HandleFunc("/cv", s.uploadCV).Methods(http.MethodPost)
	api.HandleFunc("/cv", s.getCV).Methods(http.MethodGet)

	return auth.RequireJWT(api, s.Auth, Role)
}

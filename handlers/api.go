package handlers

import (
	"net/http"

	"github.com/docsocsf/sponsor-portal/sponsor"
	"github.com/docsocsf/sponsor-portal/student"
	"github.com/gorilla/mux"
)

func Api(students *student.Service, sponsors *sponsor.Service) http.Handler {
	r := mux.NewRouter()
	r.PathPrefix("/students/").Handler(http.StripPrefix("/students", students.GetApiRoutes()))
	r.PathPrefix("/sponsors/").Handler(http.StripPrefix("/sponsors", sponsors.GetApiRoutes()))
	return r
}

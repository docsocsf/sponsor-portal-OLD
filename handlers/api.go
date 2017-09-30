package handlers

import (
	"net/http"

	"github.com/docsocsf/sponsor-portal/sponsor"
	"github.com/docsocsf/sponsor-portal/student"
	"github.com/gorilla/mux"
)

type Api struct {
	router *mux.Router
}

func NewApi(r *mux.Router, students *student.Service, sponsors *sponsor.Service) *Api {
	r.PathPrefix("/students/").Handler(http.StripPrefix("/api/students", students.GetApiRoutes()))
	r.PathPrefix("/sponsors/").Handler(http.StripPrefix("/api/sponsors", sponsors.GetApiRoutes()))
	return &Api{r}
}

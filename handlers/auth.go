package handlers

import (
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/sponsor"
	"github.com/docsocsf/sponsor-portal/student"
	"github.com/egnwd/roles"
	"github.com/gorilla/mux"
)

type Auth struct {
	router *mux.Router
}

func NewAuth(r *mux.Router, students *student.Service, sponsors *sponsor.Service) *Auth {
	r.HandleFunc("/jwt", getJwt)
	r.PathPrefix("/students/").Handler(http.StripPrefix("/auth/students", students.Auth.Handler()))
	r.PathPrefix("/sponsors/").Handler(http.StripPrefix("/auth/sponsors", sponsors.Auth.Handler()))
	return &Auth{r}
}

func getJwt(w http.ResponseWriter, r *http.Request) {
	vs, ok := r.URL.Query()["single"]
	single := ok && len(vs) > 0
	h := auth.RequireAuth(auth.GetToken(single), "/", roles.Anyone)
	h.ServeHTTP(w, r)
}

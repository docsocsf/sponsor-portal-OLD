package handlers

import (
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/sponsor"
	"github.com/docsocsf/sponsor-portal/student"
	"github.com/egnwd/roles"
	"github.com/gorilla/mux"
)

func Auth(students *student.Service, sponsors *sponsor.Service) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/jwt", getJwt)
	r.PathPrefix("/students/").Handler(http.StripPrefix("/students", students.Auth.Handler()))
	r.PathPrefix("/sponsors/").Handler(http.StripPrefix("/sponsors", sponsors.Auth.Handler()))
	return r
}

func getJwt(w http.ResponseWriter, r *http.Request) {
	vs, ok := r.URL.Query()["single"]
	single := ok && len(vs) > 0
	h := auth.RequireAuth(auth.GetToken(single), "/", roles.Anyone)
	h.ServeHTTP(w, r)
}

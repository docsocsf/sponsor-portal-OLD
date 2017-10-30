package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/cgi"

	"github.com/docsocsf/sponsor-portal/auth"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not supported", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "could not parse form", http.StatusBadRequest)
		return
	}

	username := r.PostFormValue("username")
	password := r.PostFormValue("username")

	var response struct {
		user *auth.UserInfo
		err  error `json:error`
	}

	response.user, response.err = userAuth(username, password)
	json.NewEncoder(w).Encode(response)
}

func main() {
	err := cgi.Serve(http.HandlerFunc(handler))
	if err != nil {
		log.Println(err)
	}
}

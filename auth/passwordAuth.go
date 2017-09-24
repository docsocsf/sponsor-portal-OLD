package auth

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func NewPasswordAuth(config *Config) (*PasswordAuth, error) {
	a := newAuth(config)
	auth := &a

	router := mux.NewRouter()

	router.HandleFunc(login, auth.handleLogin).Methods(http.MethodPost)
	router.HandleFunc(callback, auth.handleCallback)
	router.HandleFunc(logout, auth.handleLogout)

	auth.router = router

	return auth, nil
}

func (auth *PasswordAuth) baseUrl() string {
	return auth.baseURL
}

func (auth *PasswordAuth) session(r *http.Request, sessionKey string) (*sessions.Session, error) {
	return auth.store.Get(r, sessionKey)
}

func (auth *PasswordAuth) Handler() http.Handler {
	return auth.router
}

func (auth *PasswordAuth) handleLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	ui := UserInfo{
		&UserInfoPlus{
			Email: email,
		},
		password,
	}

	id, err := auth.get(ui)
	if err != nil {
		log.Println(err.Error())
		auth.failureHandler.ServeHTTP(w, r)
		return
	}

	err = setCurrentUser(w, r, id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth.successHandler.ServeHTTP(w, r)
}

func (auth *PasswordAuth) handleLogout(w http.ResponseWriter, r *http.Request) {
	err := deleteCurrentUser(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth.postLogoutHandler.ServeHTTP(w, r)
}

func (auth *PasswordAuth) handleCallback(w http.ResponseWriter, r *http.Request) {
	// To fullfill interface (TODO rethink if needed)
}

package auth

import (
	"log"
	"net/http"

	"github.com/docsocsf/sponsor-portal/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func NewBasicAuth(conf *Config) (*BasicAuth, error) {
	basicAuthConfig, err := config.GetBasicAuth()
	if err != nil {
		return nil, err
	}

	auth := &BasicAuth{
		auth:  newAuth(conf),
		realm: basicAuthConfig.Realm,
	}

	router := mux.NewRouter()

	router.HandleFunc(login, auth.handleLogin).Methods(http.MethodGet)
	router.HandleFunc(logout, auth.handleLogout)

	auth.router = router

	return auth, nil
}

func (auth *BasicAuth) baseUrl() string {
	return auth.baseURL
}

func (auth *BasicAuth) session(r *http.Request, sessionKey string) (*sessions.Session, error) {
	return auth.store.Get(r, sessionKey)
}

func (auth *BasicAuth) Handler() http.Handler {
	return auth.router
}

// Inspired by: https://stackoverflow.com/a/39591234
func (auth *BasicAuth) handleLogin(w http.ResponseWriter, r *http.Request) {

	user, pass, _ := r.BasicAuth()

	ui, err := userAuth(user, pass)
	if err != nil {
		log.Println(err.Error())
		w.Header().Set("WWW-Authenticate", `Basic realm="`+auth.realm+`"`)
		auth.failureHandler.ServeHTTP(w, r)
		return
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

func (auth *BasicAuth) handleLogout(w http.ResponseWriter, r *http.Request) {
	err := deleteCurrentUser(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth.postLogoutHandler.ServeHTTP(w, r)
}

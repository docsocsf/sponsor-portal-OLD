package auth

import (
	"net/http"
	"github.com/gorilla/sessions"
	"github.com/gorilla/mux"
	"github.com/docsocsf/sponsor-portal/config"
	"log"
)

func NewBasicAuth(conf *Config) (*BasicAuth, error) {
	basicAuthConfig, err := config.GetBasicAuth()
	if err != nil {
		return nil, err
	}
	auth := &BasicAuth{
		auth: newAuth(conf),
		username: basicAuthConfig.ServiceUsername,
		password: basicAuthConfig.ServicePassword,
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

	user, pass, ok := r.BasicAuth()

	wrapper := &LDAPWrapper{}

	userIsAuthenticated := wrapper.userAuth(auth.username, auth.password, user, pass)

	if  !ok || !userIsAuthenticated {
		w.Header().Set("WWW-Authenticate", `Basic realm="`+auth.realm+`"`)
		auth.failureHandler.ServeHTTP(w, r)
		return
	}

	ui := UserInfo{
		Name: wrapper.searchForName(user),
		Email: user + "ic.ac.uk",
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
package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauthService "google.golang.org/api/oauth2/v2"
)

func NewOAuth(config *Config) (*OAuth, error) {
	auth := &OAuth{
		auth: newAuth(config),
		oauth: &oauth2.Config{
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			RedirectURL:  config.BaseURL + callback,

			Endpoint: google.Endpoint,

			Scopes: scopes,
		},
	}

	router := mux.NewRouter()

	router.HandleFunc(login, auth.handleLogin)
	router.HandleFunc(callback, auth.handleCallback)
	router.HandleFunc(logout, auth.handleLogout)
	router.Handle(token, RequireAuth(auth, getToken(auth.auth)))

	auth.router = router

	return auth, nil
}

func (auth *OAuth) baseUrl() string {
	return auth.auth.baseURL
}

func (auth *OAuth) session(r *http.Request, sessionKey string) (*sessions.Session, error) {
	return auth.store.Get(r, sessionKey)
}

func (auth *OAuth) jwtConf() jwtConfig {
	return auth.jwt
}

func (auth *OAuth) Handler() http.Handler {
	return auth.router
}

func (auth *OAuth) handleLogin(w http.ResponseWriter, r *http.Request) {
	state, err := generateAndStoreState(auth, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := auth.oauth.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (auth *OAuth) handleLogout(w http.ResponseWriter, r *http.Request) {
	err := deleteCurrentUser(auth, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth.postLogoutHandler.ServeHTTP(w, r)
}

const userInfoEndpoint = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func (auth *OAuth) handleCallback(w http.ResponseWriter, r *http.Request) {
	expectedState, err := getAndDeleteState(auth, w, r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state := r.FormValue("state")
	if state != expectedState {
		http.Error(w, fmt.Sprintf("invalid oauth state, expected '%s', got '%s'",
			expectedState, state), http.StatusInternalServerError)
		return
	}

	code := r.FormValue("code")
	token, err := auth.oauth.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	src := auth.oauth.TokenSource(oauth2.NoContext, token)
	client := oauth2.NewClient(oauth2.NoContext, src)
	service, err := oauthService.New(client)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ui, err := service.Userinfo.Get().Do()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := auth.get(UserInfo{ui, ""})
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to get (or create) user", http.StatusInternalServerError)
		return
	}

	if id == nil {
		auth.failureHandler.ServeHTTP(w, r)
	}

	err = setCurrentUser(auth, w, r, id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth.successHandler.ServeHTTP(w, r)
}

package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauthService "google.golang.org/api/oauth2/v2"
)

type Auth struct {
	router *mux.Router

	store   sessions.Store
	baseURL string

	oauth *oauth2.Config

	get func(info UserInfo) (UserIdentifier, error)

	successHandler    http.Handler
	failureHandler    http.Handler
	postLogoutHandler http.Handler
}

const (
	login    = "/login"
	logout   = "/logout"
	callback = "/callback"
)

var scopes = []string{oauthService.UserinfoEmailScope, oauthService.UserinfoProfileScope}

type UserInfo *oauthService.Userinfoplus

func New(config *Config) (*Auth, error) {
	auth := &Auth{
		store:   sessions.NewCookieStore(config.CookieSecret),
		baseURL: config.BaseURL,

		oauth: &oauth2.Config{
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			RedirectURL:  config.BaseURL + callback,

			Endpoint: google.Endpoint,

			Scopes: scopes,
		},

		get: config.Get,

		successHandler:    config.SuccessHandler,
		failureHandler:    config.FailureHandler,
		postLogoutHandler: config.PostLogoutHandler,
	}

	router := mux.NewRouter()

	router.HandleFunc(login, auth.handleLogin)
	router.HandleFunc(callback, auth.handleCallback)
	router.HandleFunc(logout, auth.handleLogout)

	auth.router = router

	return auth, nil
}

func (auth *Auth) Handler() http.Handler {
	return auth.router
}

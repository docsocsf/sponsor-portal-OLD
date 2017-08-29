package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauthService "google.golang.org/api/oauth2/v2"
)

type auth struct {
	router *mux.Router

	store   sessions.Store
	baseURL string

	get func(info UserInfo) (UserIdentifier, error)

	successHandler    http.Handler
	failureHandler    http.Handler
	postLogoutHandler http.Handler

	jwt jwtConfig
}

type jwtConfig struct {
	secret []byte
	issuer string
}

type OAuth struct {
	auth
	oauth *oauth2.Config
}

const (
	login    = "/login"
	logout   = "/logout"
	callback = "/callback"
	token    = "/jwt/token"
)

var scopes = []string{oauthService.UserinfoEmailScope, oauthService.UserinfoProfileScope}

type UserInfo *oauthService.Userinfoplus

func New(config *Config) (*OAuth, error) {
	auth := &OAuth{
		auth: auth{
			store:   sessions.NewCookieStore(config.CookieSecret),
			baseURL: config.BaseURL,

			get: config.Get,

			successHandler:    config.SuccessHandler,
			failureHandler:    config.FailureHandler,
			postLogoutHandler: config.PostLogoutHandler,

			jwt: jwtConfig{
				secret: config.JwtSecret,
				issuer: config.JwtIssuer,
			},
		},
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
	router.Handle(token, auth.RequireAuth(http.HandlerFunc(auth.getToken)))

	auth.router = router

	return auth, nil
}

func (auth *OAuth) Handler() http.Handler {
	return auth.router
}

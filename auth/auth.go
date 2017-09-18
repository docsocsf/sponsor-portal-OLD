package auth

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	oauthService "google.golang.org/api/oauth2/v2"
)

type Auth interface {
	//Handlers
	Handler() http.Handler
	handleCallback(w http.ResponseWriter, r *http.Request)
	handleLogin(w http.ResponseWriter, r *http.Request)
	handleLogout(w http.ResponseWriter, r *http.Request)

	baseUrl() string
	session(r *http.Request, sessionKey string) (*sessions.Session, error)
	jwtConf() jwtConfig
}

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

type OAuth struct {
	auth
	oauth *oauth2.Config
}

type PasswordAuth = auth

var scopes = []string{oauthService.UserinfoEmailScope, oauthService.UserinfoProfileScope}

const (
	login    = "/login"
	logout   = "/logout"
	callback = "/callback"
	token    = "/jwt/token"
)

type UserInfoPlus = oauthService.Userinfoplus

type UserInfo struct {
	*UserInfoPlus
	Password string
}

func newAuth(config *Config) auth {
	return auth{
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
	}
}

func PasswordCorrect(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

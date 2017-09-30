package auth

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/egnwd/roles"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	oauthService "google.golang.org/api/oauth2/v2"
	_ "github.com/joho/godotenv/autoload"

	"github.com/docsocsf/sponsor-portal/config"
)

type Auth interface {
	//Handlers
	Handler() http.Handler
	handleLogin(w http.ResponseWriter, r *http.Request)
	handleLogout(w http.ResponseWriter, r *http.Request)

	baseUrl() string
	session(r *http.Request, sessionKey string) (*sessions.Session, error)
}

type auth struct {
	router *mux.Router

	store   *sessions.CookieStore
	baseURL string

	get func(info UserInfo) (*UserIdentifier, error)

	successHandler    http.Handler
	failureHandler    http.Handler
	postLogoutHandler http.Handler
}

type PasswordAuth = auth
type BasicAuth struct {
	auth
	username string
	password string
	realm string
}


var scopes = []string{oauthService.UserinfoEmailScope, oauthService.UserinfoProfileScope}

const (
	login    = "/login"
	logout   = "/logout"
	callback = "/callback"
)

type UserInfo struct {
	Name string
	Email string
	Password string
}

var cookieJar *sessions.CookieStore

func init() {
	cookieConfig, err := config.GetAuth()
	if err != nil {
		log.Fatal("Could not load config: ", err.Error())
	}
	cookieJar = sessions.NewCookieStore([]byte(cookieConfig.CookieSecret))
}

func newAuth(config *Config) auth {
	return auth{
		store:   cookieJar,
		baseURL: config.BaseURL,

		get: config.Get,

		successHandler:    config.SuccessHandler,
		failureHandler:    config.FailureHandler,
		postLogoutHandler: config.PostLogoutHandler,
	}
}

func PasswordCorrect(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func RoleChecker(role string) roles.Checker {
	return roles.Checker(func(req *http.Request, user interface{}) bool {
		if user != nil {
			if id, ok := user.(*UserIdentifier); ok {
				return id.Role == role
			}
		}

		return false
	})
}

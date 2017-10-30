package auth

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/egnwd/roles"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"

	"github.com/docsocsf/sponsor-portal/config"
)

type Auth interface {
	//Handlers
	Handler() http.Handler

	session(r *http.Request, sessionKey string) (*sessions.Session, error)
}

type commonAuth struct {
	router *mux.Router

	store *sessions.CookieStore

	get func(info UserInfo) (*UserIdentifier, error)

	successHandler    http.Handler
	failureHandler    http.Handler
	postLogoutHandler http.Handler
}

const (
	login  = "/login"
	logout = "/logout"
)

type UserInfo struct {
	Name     string
	Email    string
	Password string
}

var cookieJar *sessions.CookieStore

func newAuth(handlers *Config) commonAuth {
	cookieConfig, err := config.GetAuth()
	if err != nil {
		log.Fatal("Could not load config: ", err.Error())
	}

	cookieJar := sessions.NewCookieStore([]byte(cookieConfig.CookieSecret))
	return commonAuth{
		store: cookieJar,

		get: handlers.Get,

		successHandler:    handlers.SuccessHandler,
		failureHandler:    handlers.FailureHandler,
		postLogoutHandler: handlers.PostLogoutHandler,
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

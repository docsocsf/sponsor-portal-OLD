package auth

import "net/http"

type Config struct {
	Get func(info UserInfo) (*UserIdentifier, error)

	SuccessHandler    http.Handler
	FailureHandler    http.Handler
	PostLogoutHandler http.Handler
}

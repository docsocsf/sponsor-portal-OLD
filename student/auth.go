package student

import (
	"net/http"
	"net/url"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/model"
)

func (s *Service) setupAuth(config *auth.Config) (err error) {
	if config.Get == nil {
		config.Get = s.authHandler
	}

	if config.SuccessHandler == nil {
		config.SuccessHandler = http.HandlerFunc(s.authSuccessHandler)
	}
	if config.FailureHandler == nil {
		config.FailureHandler = http.HandlerFunc(s.authFailureHandler)
	}
	if config.PostLogoutHandler == nil {
		config.PostLogoutHandler = http.HandlerFunc(s.authPostLogoutHandler)
	}

	s.Auth, err = auth.New(config)

	return
}

func (s *Service) authHandler(info auth.UserInfo) (auth.UserIdentifier, error) {
	userModel := model.User{
		Name: info.Name,
		Auth: model.UserAuth{
			Email: info.Email,
		},
	}

	user, err := s.UserReader.Get(userModel)
	if err != nil {
		return nil, err
	}

	return user.Id, nil
}

func (s *Service) authSuccessHandler(w http.ResponseWriter, r *http.Request) {
	redirect(w, r, "/students")
}

func (s *Service) authFailureHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Authentication failure", http.StatusForbidden)
}

func (s *Service) authPostLogoutHandler(w http.ResponseWriter, r *http.Request) {
	redirect(w, r, "/")
}

func redirect(w http.ResponseWriter, r *http.Request, to string) {
	newUri, err := url.Parse(to)
	if err != nil {
		http.Error(w, "Failed to parse redirect path", http.StatusInternalServerError)
		return
	}

	baseUri, err := url.Parse(r.RequestURI)
	if err != nil {
		http.Error(w, "Failed to parse redirect base", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", baseUri.ResolveReference(newUri).String())
	w.WriteHeader(http.StatusSeeOther)
}

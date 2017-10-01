package sponsor

import (
	"errors"
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/httputils"
	"github.com/docsocsf/sponsor-portal/model"
	"github.com/egnwd/roles"
)

const Role = "sponsor"

func init() {
	roles.Register(Role, auth.RoleChecker(Role))
}

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

	s.Auth, err = auth.NewPasswordAuth(config)

	return
}

func (s *Service) authHandler(info auth.UserInfo) (*auth.UserIdentifier, error) {
	userModel := model.User{
		Auth: &model.UserAuth{
			Email: info.Email,
		},
	}

	hashedPassword, err := s.UserReader.HashedPassword(info.Email)
	if err != nil {
		return nil, err
	}

	if !auth.PasswordCorrect(info.Password, hashedPassword) {
		return nil, errors.New("Incorrect password")
	}

	user, err := s.UserReader.Get(userModel)
	if err != nil {
		return nil, err
	}

	id := auth.UserIdentifier{user.Id, Role}

	return &id, nil
}

func (s *Service) authSuccessHandler(w http.ResponseWriter, r *http.Request) {
	httputils.Redirect(w, r, "/sponsors/")
}

func (s *Service) authFailureHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func (s *Service) authPostLogoutHandler(w http.ResponseWriter, r *http.Request) {
	httputils.Redirect(w, r, "/")
}

package auth

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

type UserIdentifier struct {
	User int64  `json:"user,omitempty"`
	Role string `json:"role"`
}

func init() {
	gob.Register(UserIdentifier{})
}

var userSessionKey = "user"

func getCurrentUser(r *http.Request) (*UserIdentifier, error) {
	session, err := cookieJar.Get(r, sessionKey)
	if err != nil {
		return nil, err
	}

	if value, ok := session.Values[userSessionKey]; ok {
		if userId, ok := value.(UserIdentifier); ok {
			return &userId, nil
		}

		return nil, errors.New(fmt.Sprintf("Got user but was wrong type, expected: *UserIdentifier, actual: %v", reflect.ValueOf(value).Type().PkgPath()))
	}

	return nil, nil
}

func setCurrentUser(w http.ResponseWriter, r *http.Request, userID *UserIdentifier) error {
	session, err := cookieJar.Get(r, sessionKey)
	if err != nil {
		return err
	}

	session.Values[userSessionKey] = *userID

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func deleteCurrentUser(w http.ResponseWriter, r *http.Request) error {
	session, err := cookieJar.Get(r, sessionKey)
	if err != nil {
		return err
	}

	delete(session.Values, userSessionKey)

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

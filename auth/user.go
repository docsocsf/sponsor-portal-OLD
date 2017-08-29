package auth

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

type UserIdentifier interface{}

var userSessionKey = "user"

func (auth *OAuth) getCurrentUser(r *http.Request) (UserIdentifier, error) {
	session, err := auth.store.Get(r, sessionKey)
	if err != nil {
		return nil, err
	}

	if value, ok := session.Values[userSessionKey]; ok {
		if userId, ok := value.(UserIdentifier); ok {
			return userId, nil
		}

		return nil, errors.New(fmt.Sprintf("Got user but was wrong type, expected: *UserIdentifier, actual: %v", reflect.ValueOf(value).Type().PkgPath()))
	}

	return nil, nil
}

func (auth *OAuth) setCurrentUser(w http.ResponseWriter, r *http.Request, userID UserIdentifier) error {
	session, err := auth.store.Get(r, sessionKey)
	if err != nil {
		return err
	}

	session.Values[userSessionKey] = userID

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (auth *OAuth) deleteCurrentUser(w http.ResponseWriter, r *http.Request) error {
	session, err := auth.store.Get(r, sessionKey)
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

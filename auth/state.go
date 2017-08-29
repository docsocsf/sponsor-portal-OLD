package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
)

const (
	sessionKey = "auth"
	stateKey   = "state"
)

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func (auth *OAuth) generateAndStoreState(w http.ResponseWriter, r *http.Request) (string, error) {
	state := randToken()

	session, err := auth.store.Get(r, sessionKey)
	if err != nil {
		log.Println(err)
	}

	session.Values[stateKey] = state

	err = session.Save(r, w)
	if err != nil {
		log.Println("Saving")
		return "", err
	}

	return state, nil
}

func (auth *OAuth) getAndDeleteState(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := auth.store.Get(r, sessionKey)
	if err != nil {
		return "", err
	}

	state, ok := session.Values[stateKey]
	if !ok {
		return "", errors.New("Failed to get session")
	}
	delete(session.Values, stateKey)

	err = session.Save(r, w)
	return state.(string), err
}

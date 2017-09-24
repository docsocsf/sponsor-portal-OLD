package auth

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
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

func generateAndStoreState(auth Auth, w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := auth.session(r, sessionKey)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	return generateAndStore(stateKey, session, w, r)
}

func getAndDeleteState(auth Auth, w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := auth.session(r, sessionKey)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	return getAndDelete(stateKey, session, w, r)
}

func generateAndStore(key string, session *sessions.Session, w http.ResponseWriter, r *http.Request) (string, error) {
	value := randToken()

	session.Values[key] = value

	err := session.Save(r, w)
	if err != nil {
		log.Println("Saving " + key)
		return "", err
	}

	return value, nil
}

func getAndDelete(key string, session *sessions.Session, w http.ResponseWriter, r *http.Request) (string, error) {
	value, ok := session.Values[key]
	if !ok {
		return "", nil
	}
	delete(session.Values, key)

	err := session.Save(r, w)
	return value.(string), err
}

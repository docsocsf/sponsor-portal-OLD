package auth

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	oauthService "google.golang.org/api/oauth2/v2"
)

const userInfoEndpoint = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func (auth *Auth) handleCallback(w http.ResponseWriter, r *http.Request) {
	expectedState, err := auth.getAndDeleteState(w, r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state := r.FormValue("state")
	if state != expectedState {
		http.Error(w, fmt.Sprintf("invalid oauth state, expected '%s', got '%s'",
			expectedState, state), http.StatusInternalServerError)
		return
	}

	code := r.FormValue("code")
	token, err := auth.oauth.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	src := auth.oauth.TokenSource(oauth2.NoContext, token)
	client := oauth2.NewClient(oauth2.NoContext, src)
	service, err := oauthService.New(client)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ui, err := service.Userinfo.Get().Do()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := auth.get(ui)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to get (or create) user", http.StatusInternalServerError)
		return
	}

	if id == nil {
		auth.failureHandler.ServeHTTP(w, r)
	}

	err = auth.setCurrentUser(w, r, id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth.successHandler.ServeHTTP(w, r)
}

package auth

import "net/http"

func (auth *Auth) handleLogin(w http.ResponseWriter, r *http.Request) {
	state, err := auth.generateAndStoreState(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := auth.oauth.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

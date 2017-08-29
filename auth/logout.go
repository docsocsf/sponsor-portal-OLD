package auth

import (
	"net/http"
)

func (auth *OAuth) handleLogout(w http.ResponseWriter, r *http.Request) {
	err := auth.deleteCurrentUser(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth.postLogoutHandler.ServeHTTP(w, r)
}

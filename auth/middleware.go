package auth

import (
	"context"
	"log"
	"net/http"
)

type userKeyType int

const userKey userKeyType = iota

func (auth *Auth) RequireAuth(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := auth.getCurrentUser(r)
		if err != nil {
			log.Printf("Failed to get the current user from request: %v\n", r)
			http.Error(w, "Failed to get current user", http.StatusInternalServerError)
			return
		}

		if userId == nil {
			http.Redirect(w, r, auth.baseURL+"/login", http.StatusTemporaryRedirect)
			return
		}

		rWithUser := setRequestUser(r, userId)

		inner.ServeHTTP(w, rWithUser)
	})
}

func setRequestUser(r *http.Request, userId UserIdentifier) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), userKey, userId))
}

func User(r *http.Request) UserIdentifier {
	userId := r.Context().Value(userKey)

	if userId == nil {
		return nil
	}

	return userId.(UserIdentifier)
}
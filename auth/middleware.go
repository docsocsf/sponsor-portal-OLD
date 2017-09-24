package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/egnwd/roles"
)

type userKeyType int

const userKey userKeyType = iota

func setRequestUser(r *http.Request, userId *UserIdentifier) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), userKey, *userId))
}

func User(r *http.Request) *UserIdentifier {
	userId := r.Context().Value(userKey)

	if user, ok := userId.(UserIdentifier); ok {
		return &user
	}

	return nil
}

func RequireAuth(inner http.Handler, redirect string, validRoles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := getCurrentUser(r)
		if err != nil {
			log.Printf("Failed to get the current user from request: %v\n", r)
			http.Error(w, "Failed to get current user", http.StatusInternalServerError)
			return
		}

		if userId == nil {
			http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
			return
		}

		if !roles.HasRole(r, userId, validRoles...) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		rWithUser := setRequestUser(r, userId)
		inner.ServeHTTP(w, rWithUser)
	})
}

func NoAuth(inner http.Handler, redirect string, validRoles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := getCurrentUser(r)
		if err != nil {
			log.Printf("Failed to get the current user from request: %v\n", r)
			http.Error(w, "Failed to get current user", http.StatusInternalServerError)
			return
		}

		if userId == nil || !roles.HasRole(r, userId, validRoles...) {
			inner.ServeHTTP(w, r)
			return
		}

		rWithUser := setRequestUser(r, userId)
		http.Redirect(w, rWithUser, redirect, http.StatusSeeOther)
	})
}

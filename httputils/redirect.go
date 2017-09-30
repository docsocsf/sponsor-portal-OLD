package httputils

import (
	"net/http"
	"net/url"
)

func Redirect(w http.ResponseWriter, r *http.Request, to string) {
	newUri, err := url.Parse(to)
	if err != nil {
		http.Error(w, "Failed to parse redirect path", http.StatusInternalServerError)
		return
	}

	baseUri, err := url.Parse(r.RequestURI)
	if err != nil {
		http.Error(w, "Failed to parse redirect base", http.StatusInternalServerError)
		return
	}

	path := baseUri.ResolveReference(newUri).String()
	http.Redirect(w, r, path, http.StatusSeeOther)
}

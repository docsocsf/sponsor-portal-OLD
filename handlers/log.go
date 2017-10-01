package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
)

func LoggingHandler(w io.Writer, h http.Handler) http.Handler {
	return handlers.LoggingHandler(w, h)
}

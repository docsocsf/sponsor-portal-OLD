package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/docsocsf/sponsor-portal/config"
	"github.com/docsocsf/sponsor-portal/handlers"
	"github.com/docsocsf/sponsor-portal/httputils"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	host, err := config.GetHost()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter().StrictSlash(true)

	student := makeStudentService(host.StaticFiles)
	sponsor := makeSponsorService(host.StaticFiles)

	handlers.NewApi(r.PathPrefix("/api/").Subrouter(), student, sponsor)
	handlers.NewAuth(r.PathPrefix("/auth/").Subrouter(), student, sponsor)

	assets := http.FileServer(http.Dir(host.StaticFiles))
	r.PathPrefix("/assets").Handler(assets)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httputils.Redirect(w, r, "/login")
	})
	r.Handle("/login", file(host.StaticFiles, "index.html"))

	r.Handle("/students", file(host.StaticFiles, "students.html"))
	r.Handle("/sponsors", file(host.StaticFiles, "sponsors.html"))

	log.Printf("Listening on %s...", host.Port)
	log.Fatal(http.ListenAndServe(host.Port, handlers.LoggingHandler(os.Stdout, r)))
}

func rootMatcher(r *http.Request, rm *mux.RouteMatch) bool {
	return r.URL.IsAbs() && !strings.HasPrefix(r.URL.Path, "/assets")
}

func file(static, filename string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(static, filename))
	})
}

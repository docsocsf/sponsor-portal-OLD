package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/docsocsf/sponsor-portal/config"
	"github.com/gorilla/handlers"
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
	student.Handle(r)

	sponsor := makeSponsorService(host.StaticFiles)
	sponsor.Handle(r)

	assets := http.FileServer(http.Dir(host.StaticFiles))
	root := home(host.StaticFiles)

	r.PathPrefix("/assets").Handler(assets)
	r.Handle("/", root)
	r.Handle("/login", root)

	log.Printf("Listening on %s...", host.Port)
	log.Fatal(http.ListenAndServe(host.Port, handlers.LoggingHandler(os.Stdout, r)))
}

func rootMatcher(r *http.Request, rm *mux.RouteMatch) bool {
	return r.URL.IsAbs() && !strings.HasPrefix(r.URL.Path, "/assets")
}

func home(static string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(static, "index.html"))
	})
}

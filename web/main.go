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

	"github.com/docsocsf/sponsor-portal/auth"
)

func main() {

	host, err := config.GetHost()
	if err != nil {
		log.Fatal(err)
	}

	root := home(host.StaticFiles)
	r := mux.NewRouter().StrictSlash(true)

	student := makeStudentService(host.StaticFiles)
	student.Handle(r, root)

	sponsor := makeSponsorService(host.StaticFiles)
	sponsor.Handle(r, root)

	r.Handle("/jwt/token", auth.RequireAuth("/", auth.GetToken()))

	assets := http.FileServer(http.Dir(host.StaticFiles))
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

package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/config"
	"github.com/docsocsf/sponsor-portal/handlers"
	"github.com/docsocsf/sponsor-portal/httputils"
	"github.com/docsocsf/sponsor-portal/sponsor"
	"github.com/docsocsf/sponsor-portal/student"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	host, err := config.GetHost()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter().StrictSlash(true)

	studentService := makeStudentService(host.StaticFiles)
	sponsorService := makeSponsorService(host.StaticFiles)

	r.PathPrefix("/api/").Handler(http.StripPrefix("/api", handlers.Api(studentService, sponsorService)))
	r.PathPrefix("/auth/").Handler(http.StripPrefix("/auth", handlers.Auth(studentService, sponsorService)))

	assets := http.FileServer(http.Dir(host.StaticFiles))
	r.PathPrefix("/assets").Handler(assets)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httputils.Redirect(w, r, "/login")
	})

	r.Handle("/login", file(host.StaticFiles, "index.html"))
	r.Handle("/students", auth.RequireAuth(file(host.StaticFiles, "students.html"), "/auth/students/login", student.Role))
	r.Handle("/sponsors", auth.RequireAuth(file(host.StaticFiles, "sponsors.html"), "/login", sponsor.Role))

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

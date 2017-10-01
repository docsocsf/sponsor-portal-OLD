package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

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
	r.Handle("/students", auth.RequireAuth(file(host.StaticFiles, "students.html"), "/login", student.Role))
	r.Handle("/sponsors", auth.RequireAuth(file(host.StaticFiles, "sponsors.html"), "/login", sponsor.Role))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	s := http.Server{Addr: host.Port, Handler: handlers.LoggingHandler(os.Stdout, r)}

	go func() {
		log.Printf("Listening on %s...", host.Port)
		if err = s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stop
	log.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
	log.Println("Goodbye")
}

func rootMatcher(r *http.Request, rm *mux.RouteMatch) bool {
	return r.URL.IsAbs() && !strings.HasPrefix(r.URL.Path, "/assets")
}

func file(static, filename string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(static, filename))
	})
}

package main

import (
	"log"
	"net/http"

	"github.com/docsocsf/sponsor-portal/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	host, err := config.GetHost()
	if err != nil {
		log.Fatal(err)
	}

	student := makeStudentService(host.StaticFiles)
	http.Handle("/students/", http.StripPrefix("/students", student.Handler()))

	sponsor := makeSponsorService()
	http.Handle("/sponsors/", http.StripPrefix("/sponsors", sponsor.Handler()))

	root := http.FileServer(http.Dir(host.StaticFiles))
	http.Handle("/assets/", root)
	http.Handle("/", root)

	log.Printf("Listening on %s...", host.Port)
	log.Fatal(http.ListenAndServe(host.Port, nil))
}

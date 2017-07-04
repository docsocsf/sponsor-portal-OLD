package main

import (
	"log"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/config"
	"github.com/docsocsf/sponsor-portal/student"
)

func makeStudentService() *student.Service {
	authEnvConfig, err := config.GetAuth()
	if err != nil {
		log.Fatal(err, "Make student service")
	}

	authConfig := &auth.Config{
		CookieSecret: []byte(authEnvConfig.CookieSecret),

		BaseURL:      authEnvConfig.BaseURL + "/students/auth",
		Issuer:       authEnvConfig.Issuer,
		ClientID:     authEnvConfig.ClientID,
		ClientSecret: authEnvConfig.ClientSecret,
	}

	service, err := student.New(authConfig)
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	err = service.SetupDatabase(db)
	if err != nil {
		log.Fatal(err)
	}

	return service
}

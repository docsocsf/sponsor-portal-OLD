package main

import (
	"log"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/config"
	"github.com/docsocsf/sponsor-portal/student"
)

func makeStudentService(staticFiles string) *student.Service {
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

		JwtSecret: []byte(authEnvConfig.JwtSecret),
		JwtIssuer: authEnvConfig.JwtIssuer,
	}

	service, err := student.New(authConfig, staticFiles)
	if err != nil {
		log.Fatal(err)
	}

	s3, err := config.GetS3("cvs/")
	if err != nil {
		log.Fatal(err)
	}
	service.SetupStorer(s3)

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

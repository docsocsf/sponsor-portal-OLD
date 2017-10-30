package main

import (
	"log"

	"github.com/docsocsf/sponsor-portal/config"
	"github.com/docsocsf/sponsor-portal/student"
)

func makeStudentService(staticFiles string) *student.Service {
	service, err := student.New(staticFiles)
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

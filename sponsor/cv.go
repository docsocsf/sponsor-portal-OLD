package sponsor

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/docsocsf/sponsor-portal/model"
	"github.com/gorilla/mux"
)

func (s *Service) getCVs(w http.ResponseWriter, r *http.Request) {
	cvs, err := s.CVReader.GetAll()
	if err != nil {
		switch e := err.(type) {
		case model.DbError:
			if e.NotFound {
				return
			}
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	payload, _ := json.Marshal(cvs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (s *Service) downloadCV(w http.ResponseWriter, r *http.Request) {
	formId, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "No id provided", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(formId)

	cv, err := s.CVReader.Get(int64(id))
	if err != nil {
		switch e := err.(type) {
		case model.DbError:
			if e.NotFound {
				log.Println(e.Err.Error())
				http.Error(w, "CV NotFound", http.StatusNotFound)
				return
			}
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := s.s3.Get(cv.File)
	if err != nil {
		switch e := err.(type) {
		case model.DbError:
			if e.NotFound {
				http.Error(w, "CV NotFound", http.StatusNotFound)
				return
			}
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", cv.Name))
	w.Header().Set("Content-Type", *file.ContentType)
	w.Header().Set("Content-Length", strconv.Itoa(int(*file.ContentLength)))

	io.Copy(w, file.Body)
}

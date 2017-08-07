package student

import (
	"log"
	"net/http"
)

const (
	cvKey string = "cv"
)

func (s *Service) uploadCV(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile(cvKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.s3.Put(header.Filename, file)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

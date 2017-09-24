package student

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/model"
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

	h := sha256.New()
	b := bytes.NewBuffer([]byte{})
	writer := io.MultiWriter(h, b)
	_, err = io.Copy(writer, file)
	hash := hex.EncodeToString(h.Sum(nil))

	ext := path.Ext(header.Filename)
	key := hash + ext

	err = s.s3.Put(key, b)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := auth.User(r)
	cv := model.CV{Name: header.Filename, File: key}
	err = s.CVWriter.Put(*id, cv)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Service) getCV(w http.ResponseWriter, r *http.Request) {
	id := auth.User(r)
	cv, err := s.CVReader.Get(id.User)
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
	payload, _ := json.Marshal(cv)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

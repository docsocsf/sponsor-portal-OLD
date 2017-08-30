package student

import (
	"encoding/json"
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
)

func (s *Service) getUserInformation(w http.ResponseWriter, r *http.Request) {
	userId := auth.User(r)
	user, err := s.UserReader.GetById(userId)
	if err != nil {
		http.Error(w, "Could not get user", http.StatusInternalServerError)
		return
	}
	payload, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

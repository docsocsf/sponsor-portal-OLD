package student

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docsocsf/sponsor-portal/auth"
)

func (s *Service) getUserInformation(w http.ResponseWriter, r *http.Request) {
	id := auth.User(r)
	user, err := s.UserReader.GetById(*id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not get user: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	payload, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

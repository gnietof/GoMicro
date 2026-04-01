package members

import (
	"encoding/json"
	"net/http"
)

type MembersController struct {
	repository *MembersRepository
}

func NewMembersController() *MembersController {
	respository := NewMembersRepository()
	return &MembersController{repository: respository}
}

func (m *MembersController) GetMembers(w http.ResponseWriter, r *http.Request) {

	result, err := m.repository.GetMembers()
	if err != nil {
		http.Error(w, "Internal error!", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (m *MembersController) GetMemberById(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing member id!", http.StatusBadRequest)
		return
	}

	result, err := m.repository.GetMemberById(id)
	if err != nil {
		http.Error(w, "Internal error!", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

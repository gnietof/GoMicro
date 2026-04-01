package locations

import (
	"encoding/json"
	"net/http"
)

type LocationsController struct {
	repository *LocationsRepository
}

func NewLocationsController() *LocationsController {
	respository := NewLocationsRepository()
	return &LocationsController{repository: respository}
}

func (m *LocationsController) GetLocations(w http.ResponseWriter, r *http.Request) {

	result, err := m.repository.GetLocations()
	if err != nil {
		http.Error(w, "Internal error!", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (m *LocationsController) GetLocationsById(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("wlc")
	if id == "" {
		http.Error(w, "Missing location wlc!", http.StatusBadRequest)
		return
	}

	result, err := m.repository.GetLocationById(id)
	if err != nil {
		http.Error(w, "Internal error!", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

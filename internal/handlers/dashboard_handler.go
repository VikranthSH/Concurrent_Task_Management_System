package handlers

import (
	"encoding/json"
	"net/http"

	"Concurrent_Task_Management_System/internal/models"
	"Concurrent_Task_Management_System/internal/services"
)

type DashboardHandler struct {
	service *services.DashboardService
}

func NewDashboardHandler(service *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: service}
}

func (h *DashboardHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {

	currentUser := r.Context().Value("currentUser").(*models.User)

	result, err := h.service.GetDashboard(r.Context(), currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

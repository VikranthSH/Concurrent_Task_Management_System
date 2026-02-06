package handlers

import (
	"encoding/json"
	"net/http"

	"Concurrent_Task_Management_System/internal/models"
	"Concurrent_Task_Management_System/internal/services"

	"github.com/gorilla/mux"
)

type ProjectHandler struct {
	service *services.ProjectService
}

func NewProjectHandler(service *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		service: service,
	}
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdProject, err := h.service.CreateProject(r.Context(), &project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProject)
}

func (h *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	project, err := h.service.GetProjectByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.service.GetAllProjects(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch projects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) GetProjectsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["userId"]

	projects, err := h.service.GetProjectsByUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.UpdateProject(r.Context(), id, updateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := h.service.DeleteProject(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

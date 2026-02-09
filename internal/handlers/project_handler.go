package handlers

import (
	"encoding/json"
	"net/http"

	"Concurrent_Task_Management_System/internal/models"
	"Concurrent_Task_Management_System/internal/services"
	"Concurrent_Task_Management_System/internal/utils"

	"github.com/gorilla/mux"
)

type ProjectHandler struct {
	service *services.ProjectService
}

func NewProjectHandler(service *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

// =========================
// CREATE PROJECT
// =========================
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		utils.SendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	createdProject, err := h.service.CreateProject(r.Context(), &project)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusCreated,
		"Project created successfully",
		createdProject,
	)
}

// =========================
// GET PROJECT BY ID
// =========================
func (h *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	project, err := h.service.GetProjectByID(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Project fetched successfully",
		project,
	)
}

// =========================
// GET ALL PROJECTS
// =========================
func (h *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.service.GetAllProjects(r.Context())
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "failed to fetch projects")
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Projects fetched successfully",
		projects,
	)
}

// =========================
// GET PROJECTS BY USER
// =========================
func (h *ProjectHandler) GetProjectsByUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userId"]

	projects, err := h.service.GetProjectsByUser(r.Context(), userID)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Projects fetched successfully",
		projects,
	)
}

// =========================
// UPDATE PROJECT
// =========================
func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		utils.SendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.UpdateProject(r.Context(), id, updateData); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Project updated successfully",
		nil,
	)
}

// =========================
// DELETE PROJECT
// =========================
func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.service.DeleteProject(r.Context(), id); err != nil {
		utils.SendError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Project deleted successfully",
		nil,
	)
}

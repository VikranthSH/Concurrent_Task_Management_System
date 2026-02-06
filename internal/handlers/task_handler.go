package handlers

import (
	"encoding/json"
	"net/http"

	"Concurrent_Task_Management_System/internal/models"
	"Concurrent_Task_Management_System/internal/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdTask, err := h.service.CreateTask(r.Context(), &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTaskByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllTasks(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetTasksByProject(w http.ResponseWriter, r *http.Request) {
	projectIDStr := mux.Vars(r)["projectId"]

	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
	if err != nil {
		http.Error(w, "invalid project id", http.StatusBadRequest)
		return
	}

	tasks, err := h.service.GetTasksByProject(r.Context(), projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetTasksByAssignedUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := mux.Vars(r)["userId"]

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	tasks, err := h.service.GetTasksByAssignedUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetTasksByStatus(w http.ResponseWriter, r *http.Request) {
	status := mux.Vars(r)["status"]

	tasks, err := h.service.GetTasksByStatus(r.Context(), status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateTask(r.Context(), id, updateData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteTask(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

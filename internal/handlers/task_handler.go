package handlers

import (
	"encoding/json"
	"net/http"

	"Concurrent_Task_Management_System/internal/models"
	"Concurrent_Task_Management_System/internal/services"
	"Concurrent_Task_Management_System/internal/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// CREATE TASK
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		utils.SendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	createdTask, err := h.service.CreateTask(r.Context(), &task)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusCreated,
		"Task created successfully",
		createdTask,
	)
}

// GET TASK BY ID

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	task, err := h.service.GetTaskByID(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Task fetched successfully",
		task,
	)
}

// GET ALL TASKS
func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllTasks(r.Context())
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "failed to fetch tasks")
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Tasks fetched successfully",
		tasks,
	)
}

// GET TASKS BY PROJECT
func (h *TaskHandler) GetTasksByProject(w http.ResponseWriter, r *http.Request) {
	projectIDStr := mux.Vars(r)["projectId"]

	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "invalid project id")
		return
	}

	tasks, err := h.service.GetTasksByProject(r.Context(), projectID)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Tasks fetched successfully",
		tasks,
	)
}

// GET TASKS BY ASSIGNED USER
func (h *TaskHandler) GetTasksByAssignedUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := mux.Vars(r)["userId"]

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	tasks, err := h.service.GetTasksByAssignedUser(r.Context(), userID)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Tasks fetched successfully",
		tasks,
	)
}

// GET TASKS BY STATUS
func (h *TaskHandler) GetTasksByStatus(w http.ResponseWriter, r *http.Request) {
	status := mux.Vars(r)["status"]

	tasks, err := h.service.GetTasksByStatus(r.Context(), status)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Tasks fetched successfully",
		tasks,
	)
}

// UPDATE TASK
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		utils.SendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.UpdateTask(r.Context(), id, updateData); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Task updated successfully",
		nil,
	)
}

// DELETE TASK
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	if err := h.service.DeleteTask(r.Context(), id); err != nil {
		utils.SendError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Task deleted successfully",
		nil,
	)
}

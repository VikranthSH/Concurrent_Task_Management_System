package handlers

import (
	"encoding/json"
	"net/http"

	"Concurrent_Task_Management_System/internal/models"
	"Concurrent_Task_Management_System/internal/services"
	"Concurrent_Task_Management_System/internal/utils"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendError(w,http.StatusBadRequest,err.Error())
		return
	}

	createdUser, err := h.service.CreateUser(r.Context(), &user)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	utils.SendSuccess(
		w,
		http.StatusCreated,
		"User created successfully",
		createdUser,
	)

}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	user, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		utils.SendError(w,http.StatusNotFound,err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.SendSuccess(
	w,
	http.StatusCreated,
	"User created successfully",
	user,
)

}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers(r.Context())
	if err != nil {
		utils.SendError(w,http.StatusInternalServerError,err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.SendSuccess(
	w,
	http.StatusCreated,
	"User created successfully",
	users,
)

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		utils.SendError(w,http.StatusBadRequest,err.Error())

		return
	}

	err := h.service.UpdateUser(r.Context(), id, updateData)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := h.service.DeleteUser(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"Concurrent_Task_Management_System/internal/services"
	"Concurrent_Task_Management_System/internal/utils"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

type LoginRequest struct {
	UserID string `json:"user_id"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := h.userService.GetUserByUserID(r.Context(), req.UserID)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := utils.GenerateJWT(user.ID.Hex(), user.Role)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.SendSuccess(w, http.StatusOK, "Login successful", map[string]string{
		"token": token,
	})
}

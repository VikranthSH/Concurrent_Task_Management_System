package handlers

import (
	"net/http"
	"strings"

	"Concurrent_Task_Management_System/internal/services"
	"Concurrent_Task_Management_System/internal/utils"
)

type DashboardHandler struct {
	dashboardService *services.DashboardService
	userService      *services.UserService
}

func NewDashboardHandler(
	dashboardService *services.DashboardService,
	userService *services.UserService,
) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
		userService:      userService,
	}
}

func (h *DashboardHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		utils.SendError(w, http.StatusUnauthorized, "Authorization header missing")
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.ParseJWT(tokenStr)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	// ✅ JWT → DB User
	currentUser, err := h.userService.GetUserByIDFromJWT(
		r.Context(),
		claims.UserID,
	)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "User not found")
		return
	}

	result, err := h.dashboardService.GetDashboard(
		r.Context(),
		currentUser,
	)
	if err != nil {
		utils.SendError(w, http.StatusForbidden, err.Error())
		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Dashboard fetched successfully",
		result,
	)
}

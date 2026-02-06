package routes

import (
	"Concurrent_Task_Management_System/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterDashboardRoutes(router *mux.Router, handler *handlers.DashboardHandler) {
	router.HandleFunc("/dashboard", handler.GetDashboard).Methods("GET")
}

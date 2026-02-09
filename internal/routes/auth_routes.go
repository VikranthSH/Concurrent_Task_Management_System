package routes

import (
	"Concurrent_Task_Management_System/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router, handler *handlers.AuthHandler) {
	router.HandleFunc("/login", handler.Login).Methods("POST")
}

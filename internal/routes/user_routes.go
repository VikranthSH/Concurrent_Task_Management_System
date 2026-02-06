package routes

import (
	"Concurrent_Task_Management_System/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, userHandler *handlers.UserHandler) {

	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")

	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
}

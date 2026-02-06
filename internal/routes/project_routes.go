package routes

import (
	"Concurrent_Task_Management_System/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterProjectRoutes(router *mux.Router, projectHandler *handlers.ProjectHandler) {

	router.HandleFunc("/projects", projectHandler.CreateProject).Methods("POST")
	router.HandleFunc("/projects", projectHandler.GetAllProjects).Methods("GET")

	router.HandleFunc("/projects/{id}", projectHandler.GetProjectByID).Methods("GET")
	router.HandleFunc("/projects/{id}", projectHandler.UpdateProject).Methods("PUT")
	router.HandleFunc("/projects/{id}", projectHandler.DeleteProject).Methods("DELETE")

	router.HandleFunc("/users/{userId}/projects", projectHandler.GetProjectsByUser).Methods("GET")
}

package routes

import (
	"Concurrent_Task_Management_System/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(router *mux.Router, taskHandler *handlers.TaskHandler) {

	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")

	router.HandleFunc("/tasks/{id}", taskHandler.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	router.HandleFunc("/projects/{projectId}/tasks", taskHandler.GetTasksByProject).Methods("GET")
	router.HandleFunc("/users/{userId}/tasks", taskHandler.GetTasksByAssignedUser).Methods("GET")
	router.HandleFunc("/tasks/status/{status}", taskHandler.GetTasksByStatus).Methods("GET")
}

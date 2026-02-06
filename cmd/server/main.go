package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"Concurrent_Task_Management_System/internal/handlers"
	"Concurrent_Task_Management_System/internal/repositories"
	"Concurrent_Task_Management_System/internal/routes"
	"Concurrent_Task_Management_System/internal/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func main() {

	// MongoDB connection
	mongoURI := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	db := client.Database("trello_lite")

	// Initialize Repository
	userRepo := repositories.NewUserRepository(db)
	projectRepo := repositories.NewProjectRepository(db)
	taskRepo := repositories.NewTaskRepository(db)

	// Initialize Service
	userService := services.NewUserService(userRepo)
	projectService := services.NewProjectService(projectRepo)
	taskService := services.NewTaskService(taskRepo)

	// Initialize Handler
	userHandler := handlers.NewUserHandler(userService)
	projectHandler := handlers.NewProjectHandler(projectService)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Router setup
	router := mux.NewRouter()
	routes.RegisterUserRoutes(router, userHandler)
	routes.RegisterProjectRoutes(router, projectHandler)
	routes.RegisterTaskRoutes(router, taskHandler)


	//HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//shutdown
	go func() {
		log.Println("Server running on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	server.Shutdown(shutdownCtx)
	client.Disconnect(shutdownCtx)

	log.Println("Server stopped gracefully")
}

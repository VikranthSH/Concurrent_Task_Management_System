package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"Concurrent_Task_Management_System/internal/handlers"
	"Concurrent_Task_Management_System/internal/middleware"
	"Concurrent_Task_Management_System/internal/repositories"
	"Concurrent_Task_Management_System/internal/routes"
	"Concurrent_Task_Management_System/internal/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// =========================
	// MongoDB Connection
	// =========================
	mongoURI := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Mongo client error:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatal("Mongo connect error:", err)
	}

	db := client.Database("trello_lite")

	// =========================
	// Repositories
	// =========================
	userRepo := repositories.NewUserRepository(db)
	projectRepo := repositories.NewProjectRepository(db)
	taskRepo := repositories.NewTaskRepository(db)

	// =========================
	// Services
	// =========================
	projectService := services.NewProjectService(projectRepo)
	userService := services.NewUserService(userRepo, projectService)

	taskService := services.NewTaskService(taskRepo)

	dashboardService := services.NewDashboardService(
		userService,
		projectService,
		taskService,
	)

	// =========================
	// Handlers
	// =========================
	userHandler := handlers.NewUserHandler(userService)
	projectHandler := handlers.NewProjectHandler(projectService)
	taskHandler := handlers.NewTaskHandler(taskService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	// =========================
	// Router
	// =========================
	router := mux.NewRouter()

	// Auth Middleware (VERY IMPORTANT)
	router.Use(middleware.AuthMiddleware(userRepo))

	// Routes
	routes.RegisterUserRoutes(router, userHandler)
	routes.RegisterProjectRoutes(router, projectHandler)
	routes.RegisterTaskRoutes(router, taskHandler)
	routes.RegisterDashboardRoutes(router, dashboardHandler)

	// =========================
	// HTTP Server
	// =========================
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start server
	go func() {
		log.Println("Server running on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error:", err)
		}
	}()

	// =========================
	// Graceful Shutdown
	// =========================
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Println("Server shutdown error:", err)
	}

	if err := client.Disconnect(shutdownCtx); err != nil {
		log.Println("Mongo disconnect error:", err)
	}

	log.Println("✅ Server stopped gracefully")
}

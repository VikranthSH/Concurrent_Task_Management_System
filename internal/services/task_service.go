package services

import (
	"context"
	"errors"
	"time"

	"Concurrent_Task_Management_System/internal/models"
	"Concurrent_Task_Management_System/internal/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {

	if task.Title == "" {
		return nil, errors.New("title is required")
	}

	if task.ProjectID == primitive.NilObjectID {
		return nil, errors.New("projectId is required")
	}

	// validate status
	validStatus := map[string]bool{
		"Todo":        true,
		"In Progress": true,
		"Done":        true,
	}

	if task.Status != "" && !validStatus[task.Status] {
		return nil, errors.New("invalid status")
	}

	// defaults
	if task.Status == "" {
		task.Status = "Todo"
	}

	if task.Priority == "" {
		task.Priority = "Medium"
	}

	// validate due date
	if !task.DueDate.IsZero() && task.DueDate.Before(time.Now()) {
		return nil, errors.New("dueDate cannot be in the past")
	}

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetTaskByID(ctx context.Context, id primitive.ObjectID) (*models.Task, error) {
	if id == primitive.NilObjectID {
		return nil, errors.New("id is required")
	}
	return s.repo.FindByID(ctx, id)
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	return s.repo.FindAll(ctx)
}

func (s *TaskService) GetTasksByProject(ctx context.Context, projectID primitive.ObjectID) ([]models.Task, error) {
	if projectID == primitive.NilObjectID {
		return nil, errors.New("projectId is required")
	}
	return s.repo.FindByProjectID(ctx, projectID)
}

func (s *TaskService) GetTasksByAssignedUser(ctx context.Context, userID primitive.ObjectID) ([]models.Task, error) {
	if userID == primitive.NilObjectID {
		return nil, errors.New("userId is required")
	}
	return s.repo.FindByAssignedUser(ctx, userID)
}

func (s *TaskService) GetTasksByStatus(ctx context.Context, status string) ([]models.Task, error) {
	if status == "" {
		return nil, errors.New("status is required")
	}
	return s.repo.FindByStatus(ctx, status)
}

func (s *TaskService) UpdateTask(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	if id == primitive.NilObjectID {
		return errors.New("id is required")
	}

	// prevent immutable updates
	delete(update, "_id")
	delete(update, "createdAt")

	update["updatedAt"] = time.Now()
	return s.repo.UpdateByID(ctx, id, update)
}

func (s *TaskService) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	if id == primitive.NilObjectID {
		return errors.New("id is required")
	}
	return s.repo.DeleteByID(ctx, id)
}

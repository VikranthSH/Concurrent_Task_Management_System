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

type ProjectService struct {
	repo repositories.ProjectRepository
}

func NewProjectService(repo repositories.ProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}
func (s *ProjectService) CreateProject(
	ctx context.Context,
	project *models.Project,
) (*models.Project, error) {

	if project.Name == "" {
		return nil, errors.New("name is required")
	}

	if project.OwnerID == primitive.NilObjectID {
		return nil, errors.New("ownerId is required")
	}

	project.CreatedAt = time.Now()

	err := s.repo.Create(ctx, project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) GetProjectByID(ctx context.Context, id string) (*models.Project, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid project id")
	}

	return s.repo.FindByID(ctx, objID)
}

func (s *ProjectService) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProjectService) GetProjectsByUser(ctx context.Context, userIDStr string) ([]models.Project, error) {
	if userIDStr == "" {
		return nil, errors.New("userId is required")
	}

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return nil, errors.New("invalid userId")
	}

	return s.repo.FindByUserID(ctx, userID)
}

func (s *ProjectService) UpdateProject(ctx context.Context, id string, update bson.M) error {
	if id == "" {
		return errors.New("id is required")
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid project id")
	}

	return s.repo.UpdateByID(ctx, objID, update)
}

func (s *ProjectService) DeleteProject(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid project id")
	}

	return s.repo.DeleteByID(ctx, objID)
}

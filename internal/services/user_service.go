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

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	// ❌ DO NOT validate user.ID (MongoDB generates it)

	if user.UserID == "" {
		return nil, errors.New("user_id is required")
	}

	if user.Name == "" {
		return nil, errors.New("name is required")
	}

	if user.Email == "" {
		return nil, errors.New("email is required")
	}

	if user.Role == "" {
		user.Role = "user"
	}

	user.CreatedAt = time.Now()

	err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	return s.repo.FindByID(ctx, objID)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, update bson.M) error {
	if id == "" {
		return errors.New("id is required")
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user id")
	}

	return s.repo.UpdateByID(ctx, objID, update)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user id")
	}

	return s.repo.DeleteByID(ctx, objID)
}

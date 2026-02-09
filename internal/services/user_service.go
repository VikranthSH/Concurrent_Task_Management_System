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
	repo           repositories.UserRepository
	projectService *ProjectService
}

func NewUserService(
	repo repositories.UserRepository,
	projectService *ProjectService,
) *UserService {
	return &UserService{
		repo: repo,
		projectService: projectService,
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
func (s *UserService) GetUsersUnderAdmin(
	ctx context.Context,
	adminID primitive.ObjectID,
) ([]models.User, error) {

	projects, err := s.projectService.GetProjectsByOwner(ctx, adminID)
	if err != nil {
		return nil, err
	}

	userMap := make(map[primitive.ObjectID]models.User)

	for _, project := range projects {
		for _, memberID := range project.MemberIDs {
			user, err := s.repo.FindByID(ctx, memberID)
			if err == nil {
				userMap[user.ID] = *user
			}
		}
	}

	var users []models.User
	for _, u := range userMap {
		users = append(users, u)
	}

	return users, nil
}

func (s *UserService) GetUserByIDFromJWT(
	ctx context.Context,
	userID string,
) (*models.User, error) {

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	return s.GetUserByObjectID(ctx, objID)
}


func (s *UserService) GetUserByUserID(
	ctx context.Context,
	userID string,
) (*models.User, error) {
	return s.repo.FindByUserID(ctx, userID)
}
func (s *UserService) GetUserByObjectID(
	ctx context.Context,
	id primitive.ObjectID,
) (*models.User, error) {

	return s.repo.FindByID(ctx, id)
}

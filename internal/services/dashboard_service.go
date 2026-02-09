package services

import (
	"Concurrent_Task_Management_System/internal/dto"
	"Concurrent_Task_Management_System/internal/models"
	"Concurrent_Task_Management_System/internal/repositories"
	"Concurrent_Task_Management_System/internal/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DashboardService struct {
	dashboardRepo repositories.DashboardRepository
	userService   *UserService
}

func NewDashboardService(
	dashboardRepo repositories.DashboardRepository,
	userService *UserService,
) *DashboardService {
	return &DashboardService{
		dashboardRepo: dashboardRepo,
		userService:   userService,
	}
}

func (s *DashboardService) GetUserFromJWT(
	ctx context.Context,
	claims *utils.Claims,
) (*models.User, error) {

	objID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		return nil, err
	}

	return s.userService.GetUserByObjectID(ctx, objID)
}

func (s *DashboardService) GetDashboard(
	ctx context.Context,
	currentUser *models.User,
) (*dto.DashboardResponse, error) {

	response := &dto.DashboardResponse{
		CurrentUser: currentUser,
	}

	if currentUser.Role != "admin" {
		return response, nil
	}

	rawUsers, err := s.dashboardRepo.GetAdminDashboard(ctx, currentUser.ID.Hex())
	if err != nil {
		return nil, err
	}

	for _, ru := range rawUsers {

		dUser := dto.DashboardUser{
			ID:     ru["_id"].(primitive.ObjectID).Hex(),
			UserID: ru["user_id"].(string),
			Name:   ru["name"].(string),
			Role:   ru["role"].(string),
		}

		projects, _ := ru["projects"].(primitive.A)
		tasks, _ := ru["tasks"].(primitive.A)

		for _, p := range projects {

			pMap := p.(bson.M)
			projectID := pMap["_id"].(primitive.ObjectID)

			dProject := dto.DashboardProject{
				ID:   projectID.Hex(),
				Name: pMap["name"].(string),
			}

			for _, t := range tasks {
				tMap := t.(bson.M)

				if tMap["projectId"].(primitive.ObjectID) == projectID {
					dProject.Tasks = append(dProject.Tasks, dto.DashboardTask{
						Title:  tMap["title"].(string),
						Status: tMap["status"].(string),
					})
				}
			}

			dUser.Projects = append(dUser.Projects, dProject)
		}

		response.Users = append(response.Users, dUser)
	}

	return response, nil
}

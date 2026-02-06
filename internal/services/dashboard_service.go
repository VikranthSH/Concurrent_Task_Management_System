package services

import (
	"context"

	"Concurrent_Task_Management_System/internal/dto"
	"Concurrent_Task_Management_System/internal/models"
)

type DashboardService struct {
	userService    *UserService
	projectService *ProjectService
	taskService    *TaskService
}

func NewDashboardService(
	us *UserService,
	ps *ProjectService,
	ts *TaskService,
) *DashboardService {
	return &DashboardService{
		userService: us,
		projectService: ps,
		taskService: ts,
	}
}

func (s *DashboardService) GetDashboard(
	ctx context.Context,
	currentUser *models.User,
) (*dto.DashboardResponse, error) {

	response := &dto.DashboardResponse{
		CurrentUser: currentUser,
	}

	switch currentUser.Role {

	case "super_admin":
		response.Users, _ = s.userService.GetAllUsers(ctx)
		response.Projects, _ = s.projectService.GetAllProjects(ctx)
		response.Tasks, _ = s.taskService.GetAllTasks(ctx)

	case "admin":
		response.Projects, _ = s.projectService.GetProjectsByOwner(ctx, currentUser.ID)
		response.Tasks, _ = s.taskService.GetTasksByOwner(ctx, currentUser.ID)
		response.Users, _ = s.userService.GetUsersUnderAdmin(ctx, currentUser.ID)

	case "employee":
		response.Projects, _ = s.projectService.GetProjectsByMember(ctx, currentUser.ID)
		response.Tasks, _ = s.taskService.GetTasksByAssignedUser(ctx, currentUser.ID)
	}

	return response, nil
}

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
		userService:    us,
		projectService: ps,
		taskService:    ts,
	}
}

func (s *DashboardService) GetDashboard(
	ctx context.Context,
	currentUser *models.User,
) (*dto.DashboardResponse, error) {

	response := &dto.DashboardResponse{
		CurrentUser: currentUser,
		Users:       []dto.DashboardUser{},
	}

	// Fetch all data once
	users, _ := s.userService.GetAllUsers(ctx)
	projects, _ := s.projectService.GetAllProjects(ctx)
	tasks, _ := s.taskService.GetAllTasks(ctx)

	for _, u := range users {

		// =========================
		// ROLE FILTERING
		// =========================
		if currentUser.Role == "employee" && u.ID != currentUser.ID {
			continue
		}

		if currentUser.Role == "admin" {
			isUnderAdmin := false
			adminProjects, _ := s.projectService.GetProjectsByOwner(ctx, currentUser.ID)

			for _, p := range adminProjects {
				for _, m := range p.MemberIDs {
					if m == u.ID {
						isUnderAdmin = true
					}
				}
			}

			if !isUnderAdmin && u.ID != currentUser.ID {
				continue
			}
		}

		// =========================
		// BUILD DASHBOARD USER
		// =========================
		dUser := dto.DashboardUser{
			ID:        u.ID.Hex(),
			UserID:    u.UserID,
			Name:      u.Name,
			Email:     u.Email,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
			Projects:  []dto.DashboardProject{},
		}

		// =========================
		// ATTACH PROJECTS & TASKS
		// =========================
		for _, p := range projects {

			isMember := false
			for _, m := range p.MemberIDs {
				if m == u.ID {
					isMember = true
				}
			}
			if !isMember {
				continue
			}

			dProject := dto.DashboardProject{
				ID:    p.ID.Hex(),
				Name:  p.Name,
				Tasks: []dto.DashboardTask{},
			}

			for _, t := range tasks {
				if t.ProjectID == p.ID && t.AssignedTo == u.ID {
					dProject.Tasks = append(dProject.Tasks, dto.DashboardTask{
						ID:     t.ID.Hex(),
						Title:  t.Title,
						Status: t.Status,
					})
				}
			}

			dUser.Projects = append(dUser.Projects, dProject)
		}

		response.Users = append(response.Users, dUser)
	}

	return response, nil
}

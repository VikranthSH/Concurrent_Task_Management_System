package dto

import "Concurrent_Task_Management_System/internal/models"

type DashboardResponse struct {
	CurrentUser *models.User     `json:"currentUser"`
	Users       []models.User   `json:"users,omitempty"`
	Projects    []models.Project `json:"projects,omitempty"`
	Tasks       []models.Task    `json:"tasks,omitempty"`
}

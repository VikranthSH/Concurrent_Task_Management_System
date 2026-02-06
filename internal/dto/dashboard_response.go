package dto

import "Concurrent_Task_Management_System/internal/models"

type DashboardResponse struct {
	CurrentUser *models.User    `json:"currentUser"`
	Users       []DashboardUser `json:"users"`
}

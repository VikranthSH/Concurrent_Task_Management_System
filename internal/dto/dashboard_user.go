package dto

import "time"

type DashboardTask struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type DashboardProject struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Tasks []DashboardTask `json:"tasks"`
}

type DashboardUser struct {
	ID        string             `json:"id"`
	UserID    string             `json:"user_id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Role      string             `json:"role"`
	CreatedAt time.Time          `json:"createdAt"`

	Projects []DashboardProject  `json:"projects"`
}

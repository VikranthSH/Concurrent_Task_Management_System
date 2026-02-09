package dto

type APIResponse struct {
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data,omitempty"`
}

package utils

import (
	"encoding/json"
	"net/http"

	"Concurrent_Task_Management_System/internal/dto"
)

func SendSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := dto.APIResponse{
		Status:      "success",
		Description: message,
		Data:        data,
	}

	json.NewEncoder(w).Encode(resp)
}

func SendError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := dto.APIResponse{
		Status:      "fail",
		Description: message,
		Data:        nil,
	}

	json.NewEncoder(w).Encode(resp)
}

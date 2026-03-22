package model

import "github.com/google/uuid"

type CreateSPPGRequest struct {
	UserID      uuid.UUID `json:"user_id"      binding:"required"`
	NameSPPG    string    `json:"name_sppg"    binding:"required"`
	LocationURL *string   `json:"location_url"`
	Contact     *string   `json:"contact"`
}

type UpdateSPPGRequest struct {
	NameSPPG    string  `json:"name_sppg"`
	LocationURL *string `json:"location_url"`
	Contact     *string `json:"contact"`
}

type SPPGResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SPPGOK(message string, data interface{}) SPPGResponse {
	return SPPGResponse{Success: true, Message: message, Data: data}
}

func SPPGFail(message string) SPPGResponse {
	return SPPGResponse{Success: false, Message: message}
}
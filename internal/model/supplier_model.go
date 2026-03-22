package model

import "github.com/google/uuid"

type CreateSupplierRequest struct {
	UserID        uuid.UUID `json:"user_id"        binding:"required"`
	StoreName     string    `json:"store_name"     binding:"required"`
	Address       *string   `json:"address"`
	ContactNumber *string   `json:"contact_number"`
	AdminNotes    *string   `json:"admin_notes"`
}

type UpdateSupplierRequest struct {
	StoreName     string  `json:"store_name"`
	Address       *string `json:"address"`
	ContactNumber *string `json:"contact_number"`
	AdminNotes    *string `json:"admin_notes"`
}

type SupplierResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SupplierOK(message string, data interface{}) SupplierResponse {
	return SupplierResponse{Success: true, Message: message, Data: data}
}

func SupplierFail(message string) SupplierResponse {
	return SupplierResponse{Success: false, Message: message}
}
package model

import "github.com/google/uuid"

type CreateProductRequest struct {
	SupplierID  uuid.UUID `json:"supplier_id"  binding:"required"`
	ProductName string    `json:"product_name" binding:"required"`
	Category    *string   `json:"category"`
	Price       float64   `json:"price"        binding:"required,gt=0"`
	Unit        *string   `json:"unit"`
	ImageURL    *string   `json:"image_url"`
}

type UpdateProductRequest struct {
	ProductName string  `json:"product_name"`
	Category    *string `json:"category"`
	Price       float64 `json:"price"  binding:"omitempty,gt=0"`
	Unit        *string `json:"unit"`
	ImageURL    *string `json:"image_url"`
}

type ProductResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ProductOK(message string, data interface{}) ProductResponse {
	return ProductResponse{Success: true, Message: message, Data: data}
}

func ProductFail(message string) ProductResponse {
	return ProductResponse{Success: false, Message: message}
}
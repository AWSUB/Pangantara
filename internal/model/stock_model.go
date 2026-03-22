package model

import "github.com/google/uuid"

type CreateStockRequest struct {
	SupplierID    uuid.UUID `json:"supplier_id"    binding:"required"`
	ProductID     uuid.UUID `json:"product_id"     binding:"required"`
	StockQuantity int       `json:"stock_quantity" binding:"required,gte=0"`
}

type UpdateStockRequest struct {
	StockQuantity int `json:"stock_quantity" binding:"required,gte=0"`
}

type StockResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func StockOK(message string, data interface{}) StockResponse {
	return StockResponse{Success: true, Message: message, Data: data}
}

func StockFail(message string) StockResponse {
	return StockResponse{Success: false, Message: message}
}
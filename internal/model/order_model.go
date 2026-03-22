package model

import "github.com/google/uuid"

type CreateOrderRequest struct {
	SPPGID uuid.UUID         `json:"sppg_id" binding:"required"`
	Notes  *string           `json:"notes"`
	Items  []OrderDetailItem `json:"items"   binding:"required,min=1"`
}

type OrderDetailItem struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity"   binding:"required,gt=0"`
}

type UpdateOrderStatusRequest struct {
	OrderStatus string `json:"order_status" binding:"required,oneof=pending processing shipped completed cancelled"`
}

type OrderResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func OrderOK(message string, data interface{}) OrderResponse {
	return OrderResponse{Success: true, Message: message, Data: data}
}

func OrderFail(message string) OrderResponse {
	return OrderResponse{Success: false, Message: message}
}
package model

import "github.com/google/uuid"

type CreateTransactionRequest struct {
	OrderID       uuid.UUID `json:"order_id"       binding:"required"`
	PaymentMethod *string   `json:"payment_method"`
	PaymentProof  *string   `json:"payment_proof"`
	AmountPaid    float64   `json:"amount_paid"    binding:"required,gt=0"`
}

type UpdatePaymentStatusRequest struct {
	PaymentStatus string `json:"payment_status" binding:"required,oneof=unpaid waiting_confirmation paid failed"`
}

type TransactionResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func TransactionOK(message string, data interface{}) TransactionResponse {
	return TransactionResponse{Success: true, Message: message, Data: data}
}

func TransactionFail(message string) TransactionResponse {
	return TransactionResponse{Success: false, Message: message}
}

package usecase

import (
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"
	"time"

	"github.com/google/uuid"
)

func CreateTransaction(req model.CreateTransactionRequest) (*entity.Transaction, error) {
	now := time.Now()
	transaction := &entity.Transaction{
		TransactionID: uuid.New(),
		OrderID:       req.OrderID,
		PaymentMethod: req.PaymentMethod,
		PaymentStatus: entity.PaymentWaitingConfirmation,
		PaymentProof:  req.PaymentProof,
		PaymentDate:   &now,
		AmountPaid:    req.AmountPaid,
	}
	return transaction, repository.CreateTransaction(transaction)
}

func GetAllTransaction() ([]entity.Transaction, error) {
	return repository.GetAllTransaction()
}

func GetTransactionByID(id uuid.UUID) (*entity.Transaction, error) {
	return repository.GetTransactionByID(id)
}

func GetTransactionByOrderID(orderID uuid.UUID) (*entity.Transaction, error) {
	return repository.GetTransactionByOrderID(orderID)
}

func UpdateTransactionStatus(id uuid.UUID, req model.UpdatePaymentStatusRequest) error {
	return repository.UpdateTransactionStatus(id, entity.PaymentStatus(req.PaymentStatus))
}

func DeleteTransaction(id uuid.UUID) error {
	return repository.DeleteTransaction(id)
}
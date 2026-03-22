package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"

	"github.com/google/uuid"
)

func CreateTransaction(t *entity.Transaction) error {
	return postgres.DB.Create(t).Error
}

func GetTransactionByID(id uuid.UUID) (*entity.Transaction, error) {
	var t entity.Transaction
	err := postgres.DB.First(&t, "transaction_id = ?", id).Error
	return &t, err
}

func GetTransactionByOrderID(orderID uuid.UUID) (*entity.Transaction, error) {
	var t entity.Transaction
	err := postgres.DB.Where("order_id = ?", orderID).First(&t).Error
	return &t, err
}

func GetAllTransaction() ([]entity.Transaction, error) {
	var list []entity.Transaction
	err := postgres.DB.Order("created_at DESC").Find(&list).Error
	return list, err
}

func UpdateTransactionStatus(id uuid.UUID, status entity.PaymentStatus) error {
	return postgres.DB.Model(&entity.Transaction{}).
		Where("transaction_id = ?", id).
		Update("payment_status", status).Error
}

func DeleteTransaction(id uuid.UUID) error {
	return postgres.DB.Delete(&entity.Transaction{}, "transaction_id = ?", id).Error
}
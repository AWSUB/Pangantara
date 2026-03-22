package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"

	"github.com/google/uuid"
)

func CreateOrder(o *entity.Order) error {
	return postgres.DB.Create(o).Error
}

func GetAllOrder() ([]entity.Order, error) {
	var list []entity.Order
	err := postgres.DB.Order("order_date DESC").Find(&list).Error
	return list, err
}

func GetOrderByID(id uuid.UUID) (*entity.Order, error) {
	var o entity.Order
	err := postgres.DB.Preload("OrderDetail").Preload("Transaction").
		First(&o, "order_id = ?", id).Error
	return &o, err
}

func GetOrderBySPPGID(sppgID uuid.UUID) ([]entity.Order, error) {
	var list []entity.Order
	err := postgres.DB.Where("sppg_id = ?", sppgID).
		Order("order_date DESC").Find(&list).Error
	return list, err
}

func GetOrderByStatus(status entity.OrderStatus) ([]entity.Order, error) {
	var list []entity.Order
	err := postgres.DB.Where("order_status = ?", status).
		Order("order_date DESC").Find(&list).Error
	return list, err
}

func UpdateOrderStatus(id uuid.UUID, status entity.OrderStatus) error {
	return postgres.DB.Model(&entity.Order{}).
		Where("order_id = ?", id).
		Update("order_status", status).Error
}

func DeleteOrder(id uuid.UUID) error {
	return postgres.DB.Delete(&entity.Order{}, "order_id = ?", id).Error
}
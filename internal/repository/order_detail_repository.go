package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"

	"github.com/google/uuid"
)

func CreateOrderDetail(d *entity.OrderDetail) error {
	return postgres.DB.Create(d).Error
}

func GetOrderDetailByOrderID(orderID uuid.UUID) ([]entity.OrderDetail, error) {
	var list []entity.OrderDetail
	err := postgres.DB.Preload("Product").
		Where("order_id = ?", orderID).Find(&list).Error
	return list, err
}

func GetOrderDetailByID(id uuid.UUID) (*entity.OrderDetail, error) {
	var d entity.OrderDetail
	err := postgres.DB.Preload("Product").
		First(&d, "detail_id = ?", id).Error
	return &d, err
}

func UpdateOrderDetail(id uuid.UUID, data map[string]interface{}) error {
	return postgres.DB.Model(&entity.OrderDetail{}).
		Where("detail_id = ?", id).Updates(data).Error
}

func DeleteOrderDetail(id uuid.UUID) error {
	return postgres.DB.Delete(&entity.OrderDetail{}, "detail_id = ?", id).Error
}
package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"

	"github.com/google/uuid"
)

func CreateStock(s *entity.Stock) error {
	return postgres.DB.Create(s).Error
}

func GetAllStock() ([]entity.Stock, error) {
	var list []entity.Stock
	err := postgres.DB.Find(&list).Error
	return list, err
}

func GetStockByID(id uuid.UUID) (*entity.Stock, error) {
	var s entity.Stock
	err := postgres.DB.First(&s, "stock_id = ?", id).Error
	return &s, err
}

func GetStockByProductID(productID uuid.UUID) (*entity.Stock, error) {
	var s entity.Stock
	err := postgres.DB.Where("product_id = ?", productID).First(&s).Error
	return &s, err
}

func GetStockBySupplierID(supplierID uuid.UUID) ([]entity.Stock, error) {
	var list []entity.Stock
	err := postgres.DB.Where("supplier_id = ?", supplierID).Find(&list).Error
	return list, err
}

func UpdateStockQuantity(productID uuid.UUID, quantity int) error {
	return postgres.DB.Model(&entity.Stock{}).
		Where("product_id = ?", productID).
		Update("stock_quantity", quantity).Error
}

func DeleteStock(id uuid.UUID) error {
	return postgres.DB.Delete(&entity.Stock{}, "stock_id = ?", id).Error
}
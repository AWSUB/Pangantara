package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"

	"github.com/google/uuid"
)

func CreateProduct(p *entity.Product) error {
	return postgres.DB.Create(p).Error
}

func GetAllProduct() ([]entity.Product, error) {
	var list []entity.Product
	err := postgres.DB.Order("product_name ASC").Find(&list).Error
	return list, err
}

func GetProductByID(id uuid.UUID) (*entity.Product, error) {
	var p entity.Product
	err := postgres.DB.First(&p, "product_id = ?", id).Error
	return &p, err
}

func GetProductBySupplier(supplierID uuid.UUID) ([]entity.Product, error) {
	var list []entity.Product
	err := postgres.DB.Where("supplier_id = ?", supplierID).
		Order("product_name ASC").Find(&list).Error
	return list, err
}

func GetProductByCategory(category string) ([]entity.Product, error) {
	var list []entity.Product
	err := postgres.DB.Where("category = ?", category).
		Order("product_name ASC").Find(&list).Error
	return list, err
}

func GetProductBySupplierAndCategory(supplierID uuid.UUID, category string) ([]entity.Product, error) {
	var list []entity.Product
	err := postgres.DB.Where("supplier_id = ? AND category = ?", supplierID, category).
		Order("product_name ASC").Find(&list).Error
	return list, err
}

func UpdateProduct(id uuid.UUID, data map[string]interface{}) error {
	return postgres.DB.Model(&entity.Product{}).Where("product_id = ?", id).Updates(data).Error
}

func DeleteProduct(id uuid.UUID) error {
	return postgres.DB.Delete(&entity.Product{}, "product_id = ?", id).Error
}
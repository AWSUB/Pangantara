package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"

	"github.com/google/uuid"
)

func CreateSupplier(s *entity.Supplier) error {
	return postgres.DB.Create(s).Error
}

func GetAllSupplier() ([]entity.Supplier, error) {
	var list []entity.Supplier
	err := postgres.DB.Order("store_name ASC").Find(&list).Error
	return list, err
}

func GetSupplierByID(id uuid.UUID) (*entity.Supplier, error) {
	var s entity.Supplier
	err := postgres.DB.First(&s, "supplier_id = ?", id).Error
	return &s, err
}

func GetSupplierByUserID(userID uuid.UUID) (*entity.Supplier, error) {
	var s entity.Supplier
	err := postgres.DB.Where("user_id = ?", userID).First(&s).Error
	return &s, err
}

func UpdateSupplier(id uuid.UUID, data map[string]interface{}) error {
	return postgres.DB.Model(&entity.Supplier{}).Where("supplier_id = ?", id).Updates(data).Error
}

func DeleteSupplier(id uuid.UUID) error {
	return postgres.DB.Delete(&entity.Supplier{}, "supplier_id = ?", id).Error
}
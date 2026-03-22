package usecase

import (
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"

	"github.com/google/uuid"
)

func CreateSupplier(req model.CreateSupplierRequest) (*entity.Supplier, error) {
	supplier := &entity.Supplier{
		SupplierID:    uuid.New(),
		UserID:        req.UserID,
		StoreName:     req.StoreName,
		Address:       req.Address,
		ContactNumber: req.ContactNumber,
		AdminNotes:    req.AdminNotes,
	}
	return supplier, repository.CreateSupplier(supplier)
}

func GetAllSupplier() ([]entity.Supplier, error) {
	return repository.GetAllSupplier()
}

func GetSupplierByID(id uuid.UUID) (*entity.Supplier, error) {
	return repository.GetSupplierByID(id)
}

func GetSupplierByUserID(userID uuid.UUID) (*entity.Supplier, error) {
	return repository.GetSupplierByUserID(userID)
}

func UpdateSupplier(id uuid.UUID, req model.UpdateSupplierRequest) error {
	data := map[string]interface{}{}
	if req.StoreName != "" {
		data["store_name"] = req.StoreName
	}
	if req.Address != nil {
		data["address"] = req.Address
	}
	if req.ContactNumber != nil {
		data["contact_number"] = req.ContactNumber
	}
	if req.AdminNotes != nil {
		data["admin_notes"] = req.AdminNotes
	}
	return repository.UpdateSupplier(id, data)
}

func DeleteSupplier(id uuid.UUID) error {
	return repository.DeleteSupplier(id)
}
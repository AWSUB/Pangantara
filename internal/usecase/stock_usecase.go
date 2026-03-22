package usecase

import (
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"

	"github.com/google/uuid"
)

func CreateStock(req model.CreateStockRequest) (*entity.Stock, error) {
	stock := &entity.Stock{
		StockID:       uuid.New(),
		SupplierID:    req.SupplierID,
		ProductID:     req.ProductID,
		StockQuantity: req.StockQuantity,
	}
	return stock, repository.CreateStock(stock)
}

func GetAllStock() ([]entity.Stock, error) {
	return repository.GetAllStock()
}

func GetStockByID(id uuid.UUID) (*entity.Stock, error) {
	return repository.GetStockByID(id)
}

func GetStockByProductID(productID uuid.UUID) (*entity.Stock, error) {
	return repository.GetStockByProductID(productID)
}

func GetStockBySupplierID(supplierID uuid.UUID) ([]entity.Stock, error) {
	return repository.GetStockBySupplierID(supplierID)
}

func UpdateStockQuantity(productID uuid.UUID, req model.UpdateStockRequest) error {
	return repository.UpdateStockQuantity(productID, req.StockQuantity)
}

func DeleteStock(id uuid.UUID) error {
	return repository.DeleteStock(id)
}
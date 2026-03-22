package usecase

import (
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"

	"github.com/google/uuid"
)

func CreateProduct(req model.CreateProductRequest) (*entity.Product, error) {
	product := &entity.Product{
		ProductID:   uuid.New(),
		SupplierID:  req.SupplierID,
		ProductName: req.ProductName,
		Category:    req.Category,
		Price:       req.Price,
		Unit:        req.Unit,
		ImageURL:    req.ImageURL,
	}
	return product, repository.CreateProduct(product)
}

func GetAllProduct() ([]entity.Product, error) {
	return repository.GetAllProduct()
}

func GetProductByID(id uuid.UUID) (*entity.Product, error) {
	return repository.GetProductByID(id)
}

func GetProductBySupplier(supplierID uuid.UUID) ([]entity.Product, error) {
	return repository.GetProductBySupplier(supplierID)
}

func GetProductByCategory(category string) ([]entity.Product, error) {
	return repository.GetProductByCategory(category)
}

func GetProductBySupplierAndCategory(supplierID uuid.UUID, category string) ([]entity.Product, error) {
	return repository.GetProductBySupplierAndCategory(supplierID, category)
}

func UpdateProduct(id uuid.UUID, req model.UpdateProductRequest) error {
	data := map[string]interface{}{}
	if req.ProductName != "" {
		data["product_name"] = req.ProductName
	}
	if req.Category != nil {
		data["category"] = req.Category
	}
	if req.Price > 0 {
		data["price"] = req.Price
	}
	if req.Unit != nil {
		data["unit"] = req.Unit
	}
	if req.ImageURL != nil {
		data["image_url"] = req.ImageURL
	}
	return repository.UpdateProduct(id, data)
}

func DeleteProduct(id uuid.UUID) error {
	return repository.DeleteProduct(id)
}
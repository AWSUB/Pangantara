package usecase

import (
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"
	"time"

	"github.com/google/uuid"
)

func CreateOrder(req model.CreateOrderRequest) (*entity.Order, error) {
	orderID := uuid.New()
	var totalAmount float64

	// Hitung total amount dari semua item
	for _, item := range req.Items {
		product, err := repository.GetProductByID(item.ProductID)
		if err != nil {
			return nil, err
		}
		totalAmount += product.Price * float64(item.Quantity)
	}

	order := &entity.Order{
		OrderID:     orderID,
		SPPGID:      req.SPPGID,
		OrderDate:   time.Now(),
		OrderStatus: entity.OrderPending,
		TotalAmount: totalAmount,
		Notes:       req.Notes,
	}

	if err := repository.CreateOrder(order); err != nil {
		return nil, err
	}

	// Simpan setiap item order detail
	for _, item := range req.Items {
		detail := &entity.OrderDetail{
			DetailID:  uuid.New(),
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		if err := repository.CreateOrderDetail(detail); err != nil {
			return nil, err
		}
	}

	return order, nil
}

func GetAllOrder() ([]entity.Order, error) {
	return repository.GetAllOrder()
}

func GetOrderByID(id uuid.UUID) (*entity.Order, error) {
	return repository.GetOrderByID(id)
}

func GetOrderBySPPGID(sppgID uuid.UUID) ([]entity.Order, error) {
	return repository.GetOrderBySPPGID(sppgID)
}

func GetOrderByStatus(status entity.OrderStatus) ([]entity.Order, error) {
	return repository.GetOrderByStatus(status)
}

func UpdateOrderStatus(id uuid.UUID, req model.UpdateOrderStatusRequest) error {
	return repository.UpdateOrderStatus(id, entity.OrderStatus(req.OrderStatus))
}

func DeleteOrder(id uuid.UUID) error {
	return repository.DeleteOrder(id)
}
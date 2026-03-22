package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Base
	OrderID     uuid.UUID   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"order_id"`
	SPPGID      uuid.UUID   `gorm:"type:uuid;not null"                              json:"sppg_id"`
	OrderDate   time.Time   `gorm:"not null;default:now()"                          json:"order_date"`
	OrderStatus OrderStatus `gorm:"type:varchar(30);not null;default:'pending'"     json:"order_status"`
	TotalAmount float64     `gorm:"type:decimal(14,2);not null;default:0"           json:"total_amount"`
	Notes       *string     `gorm:"type:text"                                       json:"notes,omitempty"`
}

func (Order) TableName() string { 
	return "orders" 
}
package entity

import "github.com/google/uuid"

type OrderDetail struct {
	DetailID  uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"detail_id"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null"                              json:"order_id"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"                              json:"product_id"`
	Quantity  int       `gorm:"type:int;not null;default:1"                     json:"quantity"`
}

func (OrderDetail) TableName() string { 
	return "order_details"
}
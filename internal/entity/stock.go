package entity

import "github.com/google/uuid"

type Stock struct {
	Base
	StockID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"stock_id"`
	SupplierID    uuid.UUID `gorm:"type:uuid;not null"                              json:"supplier_id"`
	ProductID     uuid.UUID `gorm:"type:uuid;not null"                              json:"product_id"`
	StockQuantity int       `gorm:"type:int;not null;default:0"                     json:"stock_quantity"`
}

func (Stock) TableName() string { 
	return "stocks"
}
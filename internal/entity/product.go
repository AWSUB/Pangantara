package entity

import "github.com/google/uuid"

type Product struct {
	Base
	ProductID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"product_id"`
	SupplierID  uuid.UUID `gorm:"type:uuid;not null"                              json:"supplier_id"`
	ProductName string    `gorm:"type:varchar(150);not null"                      json:"product_name"`
	Category    *string   `gorm:"type:varchar(50)"                                json:"category,omitempty"`
	Price       float64   `gorm:"type:decimal(12,2);not null;default:0"           json:"price"`
	Unit        *string   `gorm:"type:varchar(50)"                                json:"unit,omitempty"`
	ImageURL    *string   `gorm:"type:varchar(255)"                               json:"image_url,omitempty"`
}

func (Product) TableName() string { 
	return "products" 
}
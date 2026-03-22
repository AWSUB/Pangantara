package entity

import "github.com/google/uuid"

type Supplier struct {
	Base
	SupplierID    uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"supplier_id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null"                              json:"user_id"`
	StoreName     string    `gorm:"type:varchar(150);not null"                      json:"store_name"`
	Address       *string   `gorm:"type:varchar"                                    json:"address,omitempty"`
	ContactNumber *string   `gorm:"type:varchar(20)"                                json:"contact_number,omitempty"`
	AdminNotes    *string   `gorm:"type:text"                                       json:"admin_notes,omitempty"`
}

func (Supplier) TableName() string {
	return "suppliers" 
}
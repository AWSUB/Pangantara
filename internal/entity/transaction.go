package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Base
	TransactionID uuid.UUID     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"transaction_id"`
	OrderID       uuid.UUID     `gorm:"type:uuid;not null"                              json:"order_id"`
	PaymentMethod *string       `gorm:"type:varchar(50)"                                json:"payment_method,omitempty"`
	PaymentStatus PaymentStatus `gorm:"type:varchar(30);not null;default:'unpaid'"      json:"payment_status"`
	PaymentProof  *string       `gorm:"type:varchar(255)"                               json:"payment_proof,omitempty"`
	PaymentDate   *time.Time    `gorm:"type:timestamp"                                  json:"payment_date,omitempty"`
	AmountPaid    float64       `gorm:"type:decimal(14,2);not null;default:0"           json:"amount_paid"`
}

func (Transaction) TableName() string { 
	return "transactions" 
}
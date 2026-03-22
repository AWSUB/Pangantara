package entity

import "github.com/google/uuid"

type User struct {
	Base
	UserID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"user_id"`
	Name     string    `gorm:"type:varchar(100);not null"                      json:"name"`
	Email    string    `gorm:"type:varchar(100);not null;uniqueIndex"          json:"email"`
	Password string    `gorm:"type:varchar(250);not null"                      json:"-"`
	Role     UserRole  `gorm:"type:varchar(20);not null;default:'sppg'"        json:"role"`
}

func (User) TableName() string { 
	return "users" 
}
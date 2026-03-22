package entity

import "github.com/google/uuid"

type SPPG struct {
	Base
	SPPGID      uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"sppg_id"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"                              json:"user_id"`
	NameSPPG    string    `gorm:"type:varchar(150);not null"                      json:"name_sppg"`
	LocationURL *string   `gorm:"type:varchar"                                    json:"location_url,omitempty"`
	Contact     *string   `gorm:"type:varchar(20)"                                json:"contact,omitempty"`
}

func (SPPG) TableName() string { 
	return "sppg" 
}
package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"

	"github.com/google/uuid"
)

func CreateSPPG(s *entity.SPPG) error {
	return postgres.DB.Create(s).Error
}

func GetAllSPPG() ([]entity.SPPG, error) {
	var list []entity.SPPG
	err := postgres.DB.Order("name_sppg ASC").Find(&list).Error
	return list, err
}

func GetSPPGByID(id uuid.UUID) (*entity.SPPG, error) {
	var s entity.SPPG
	err := postgres.DB.First(&s, "sppg_id = ?", id).Error
	return &s, err
}

func GetSPPGByUserID(userID uuid.UUID) ([]entity.SPPG, error) {
	var list []entity.SPPG
	err := postgres.DB.Where("user_id = ?", userID).Order("name_sppg ASC").Find(&list).Error
	return list, err
}

func UpdateSPPG(id uuid.UUID, data map[string]interface{}) error {
	return postgres.DB.Model(&entity.SPPG{}).Where("sppg_id = ?", id).Updates(data).Error
}

func DeleteSPPG(id uuid.UUID) error {
	return postgres.DB.Delete(&entity.SPPG{}, "sppg_id = ?", id).Error
}
package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"

	"github.com/google/uuid"
)

func CreateUser(u *entity.User) error {
	return postgres.DB.Create(u).Error
}

func GetAllUser() ([]entity.User, error) {
	var list []entity.User
	err := postgres.DB.Order("name ASC").Find(&list).Error
	return list, err
}

func GetUserByID(id uuid.UUID) (*entity.User, error) {
	var u entity.User
	err := postgres.DB.First(&u, "user_id = ?", id).Error
	return &u, err
}

func GetUserByEmail(email string) (*entity.User, error) {
	var u entity.User
	err := postgres.DB.Where("email = ?", email).First(&u).Error
	return &u, err
}

func UpdateUser(id uuid.UUID, data map[string]interface{}) error {
	return postgres.DB.Model(&entity.User{}).Where("user_id = ?", id).Updates(data).Error
}

func DeleteUser(id uuid.UUID) error {
	return postgres.DB.Delete(&entity.User{}, "user_id = ?", id).Error
}
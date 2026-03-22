package usecase

import (
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"

	"github.com/google/uuid"
)

func CreateUser(req model.CreateUserRequest) (*entity.User, error) {
	user := &entity.User{
		UserID:   uuid.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, // TODO: bcrypt hash
		Role:     entity.UserRole(req.Role),
	}
	return user, repository.CreateUser(user)
}

func GetAllUser() ([]entity.User, error) {
	return repository.GetAllUser()
}

func GetUserByID(id uuid.UUID) (*entity.User, error) {
	return repository.GetUserByID(id)
}

func UpdateUser(id uuid.UUID, req model.UpdateUserRequest) error {
	data := map[string]interface{}{}
	if req.Name != "" {
		data["name"] = req.Name
	}
	if req.Email != "" {
		data["email"] = req.Email
	}
	return repository.UpdateUser(id, data)
}

func DeleteUser(id uuid.UUID) error {
	return repository.DeleteUser(id)
}
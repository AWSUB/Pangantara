package usecase

import (
	"errors"
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"
	"sppg-backend/pkg/jwt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(req model.CreateUserRequest) (*entity.User, error) {
	// Cek apakah email sudah dipakai
	existing, err := repository.GetUserByEmail(req.Email)
	if err == nil && existing != nil {
		return nil, errors.New("email sudah digunakan")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		UserID:   uuid.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     entity.UserRole(req.Role),
	}

	return user, repository.CreateUser(user)
}

func Login(req model.LoginRequest) (*model.LoginResponse, error) {
	// Cari user berdasarkan email
	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("email atau password salah")
	}

	// Generate token
	token, err := jwt.GenerateToken(user.UserID.String(), user.Email, string(user.Role))
	if err != nil {
		return nil, err
	}

	response := model.LoginOK("Login berhasil", token, map[string]interface{}{
		"user_id": user.UserID,
		"name":    user.Name,
		"email":   user.Email,
		"role":    user.Role,
	})

	return &response, nil
}
package usecase

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"
	"sppg-backend/pkg/email"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ForgotPassword(req model.ForgotPasswordRequest) error {
	// Cek apakah email terdaftar
	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		// Jangan beritahu user kalau email tidak ada (security)
		return nil
	}

	// Hapus token lama yang expired
	repository.DeleteExpiredResetPassword(user.UserID)

	// Generate token random
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return err
	}
	token := hex.EncodeToString(tokenBytes)

	// Simpan token ke database
	resetPassword := &entity.ResetPassword{
		ID:        uuid.New(),
		UserID:    user.UserID,
		Token:     token,
		ExpiredAt: time.Now().Add(15 * time.Minute),
		IsUsed:    false,
	}

	if err := repository.CreateResetPassword(resetPassword); err != nil {
		return err
	}

	// Kirim email
	resetLink := "https://pangantara.vercel.app/reset-password?token=" + token
	if err := email.SendForgotPasswordEmail(user.Email, user.Name, resetLink); err != nil {
		return err
	}

	return nil
}

func ResetPassword(req model.ResetPasswordRequest) error {
	// Cek apakah password dan konfirmasi password sama
	if req.NewPassword != req.ConfirmPassword {
		return errors.New("kata sandi tidak cocok")
	}

	// Cek token
	resetPassword, err := repository.GetResetPasswordByToken(req.Token)
	if err != nil {
		return errors.New("token tidak valid atau sudah kadaluarsa")
	}

	// Hash password baru
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password user
	if err := repository.UpdateUser(resetPassword.UserID, map[string]interface{}{
		"password": string(hashedPassword),
	}); err != nil {
		return err
	}

	// Tandai token sebagai sudah dipakai
	if err := repository.MarkResetPasswordAsUsed(resetPassword.ID); err != nil {
		return err
	}

	return nil
}
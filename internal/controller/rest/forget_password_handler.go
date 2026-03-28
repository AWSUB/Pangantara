package rest

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func ForgotPasswordRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/forgot-password", forgotPassword)
		auth.POST("/reset-password", resetPassword)
	}
}

func forgotPassword(c *gin.Context) {
	var req model.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	if err := usecase.ForgotPassword(req); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OKMessage("If your email is registered, a reset link will be sent", nil))
}

func resetPassword(c *gin.Context) {
	var req model.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	if err := usecase.ResetPassword(req); err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.OKMessage("Password reset successfully", nil))
}
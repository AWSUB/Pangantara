package rest

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", register)
		auth.POST("/login", login)
		auth.POST("/refresh", refreshToken)
	}
}

func register(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	data, err := usecase.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, model.Created(data))
}

func login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	response, err := usecase.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.BadRequest(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response)
}

func refreshToken(c *gin.Context) {
	var req model.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	response, err := usecase.RefreshToken(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.BadRequest(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response)
}
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
	}
}

func register(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.LoginFail(err.Error()))
		return
	}
	data, err := usecase.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.LoginFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, model.LoginOK("Register berhasil", "", data))
}

func login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.LoginFail(err.Error()))
		return
	}
	response, err := usecase.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.LoginFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response)
}
package rest

import (
	"net/http"
	"sppg-backend/internal/middleware"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		// Hanya admin yang bisa akses
		users.POST("", middleware.RoleMiddleware("admin"), createUser)
		users.GET("", middleware.RoleMiddleware("admin"), getAllUser)
		users.GET("/:id", middleware.RoleMiddleware("admin", "supplier", "sppg"), getUserByID)
		users.PUT("/:id", middleware.RoleMiddleware("admin", "supplier", "sppg"), updateUser)
		users.DELETE("/:id", middleware.RoleMiddleware("admin"), deleteUser)
	}
}

func createUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.UserFail(err.Error()))
		return
	}
	data, err := usecase.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.UserFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, model.UserOK("User berhasil dibuat", data))
}

func getAllUser(c *gin.Context) {
	list, err := usecase.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.UserFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.UserOK("OK", list))
}

func getUserByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.UserFail("ID tidak valid"))
		return
	}
	data, err := usecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.UserFail("User tidak ditemukan"))
		return
	}
	c.JSON(http.StatusOK, model.UserOK("OK", data))
}

func updateUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.UserFail("ID tidak valid"))
		return
	}
	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.UserFail(err.Error()))
		return
	}
	if err := usecase.UpdateUser(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.UserFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.UserOK("User berhasil diupdate", nil))
}

func deleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.UserFail("ID tidak valid"))
		return
	}
	if err := usecase.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.UserFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.UserOK("User berhasil dihapus", nil))
}
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
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	data, err := usecase.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusCreated, model.Created(data))
}

func getAllUser(c *gin.Context) {
	list, err := usecase.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OK(list))
}

func getUserByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	data, err := usecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NotFound("User"))
		return
	}
	c.JSON(http.StatusOK, model.OK(data))
}

func updateUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	if err := usecase.UpdateUser(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.Updated())
}

func deleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	if err := usecase.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.Deleted())
}
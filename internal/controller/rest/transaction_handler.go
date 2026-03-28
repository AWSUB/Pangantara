package rest

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"
	"sppg-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TransactionRoutes(r *gin.RouterGroup) {
	transaction := r.Group("/transactions")
	{
		transaction.POST("", createTransaction)
		transaction.GET("", getAllTransaction)
		transaction.GET("/:id", getTransactionByID)
		transaction.GET("/order/:order_id", getTransactionByOrderID)
		transaction.PUT("/:id/status", middleware.RoleMiddleware("admin") , updateTransactionStatus)
		transaction.DELETE("/:id", middleware.RoleMiddleware("admin") , deleteTransaction)
	}
}

func createTransaction(c *gin.Context) {
	var req model.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	data, err := usecase.CreateTransaction(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusCreated, model.Created(data))
}

func getAllTransaction(c *gin.Context) {
	list, err := usecase.GetAllTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OK(list))
}

func getTransactionByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	data, err := usecase.GetTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NotFound("Transaction"))
		return
	}
	c.JSON(http.StatusOK, model.OK(data))
}

func getTransactionByOrderID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	data, err := usecase.GetTransactionByOrderID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NotFound("Transaction"))
		return
	}
	c.JSON(http.StatusOK, model.OK(data))
}

func updateTransactionStatus(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	var req model.UpdatePaymentStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	if err := usecase.UpdateTransactionStatus(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OKMessage("Payment status updated successfully", nil))
}

func deleteTransaction(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	if err := usecase.DeleteTransaction(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.Deleted())
}
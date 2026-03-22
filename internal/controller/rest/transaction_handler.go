package rest

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

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
		transaction.PUT("/:id/status", updateTransactionStatus)
		transaction.DELETE("/:id", deleteTransaction)
	}
}

func createTransaction(c *gin.Context) {
	var req model.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.TransactionFail(err.Error()))
		return
	}
	data, err := usecase.CreateTransaction(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.TransactionFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, model.TransactionOK("Transaksi berhasil dibuat", data))
}

func getAllTransaction(c *gin.Context) {
	list, err := usecase.GetAllTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.TransactionFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.TransactionOK("OK", list))
}

func getTransactionByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.TransactionFail("ID tidak valid"))
		return
	}
	data, err := usecase.GetTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.TransactionFail("Transaksi tidak ditemukan"))
		return
	}
	c.JSON(http.StatusOK, model.TransactionOK("OK", data))
}

func getTransactionByOrderID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.TransactionFail("ID tidak valid"))
		return
	}
	data, err := usecase.GetTransactionByOrderID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.TransactionFail("Transaksi tidak ditemukan"))
		return
	}
	c.JSON(http.StatusOK, model.TransactionOK("OK", data))
}

func updateTransactionStatus(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.TransactionFail("ID tidak valid"))
		return
	}
	var req model.UpdatePaymentStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.TransactionFail(err.Error()))
		return
	}
	if err := usecase.UpdateTransactionStatus(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.TransactionFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.TransactionOK("Status transaksi berhasil diupdate", nil))
}

func deleteTransaction(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.TransactionFail("ID tidak valid"))
		return
	}
	if err := usecase.DeleteTransaction(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.TransactionFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.TransactionOK("Transaksi berhasil dihapus", nil))
}
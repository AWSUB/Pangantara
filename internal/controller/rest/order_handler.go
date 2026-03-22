package rest

import (
	"net/http"
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func OrderRoutes(r *gin.RouterGroup) {
	order := r.Group("/orders")
	{
		order.POST("", createOrder)
		order.GET("", getAllOrder)
		order.GET("/:id", getOrderByID)
		order.GET("/sppg/:sppg_id", getOrderBySPPGID)
		order.PUT("/:id/status", updateOrderStatus)
		order.DELETE("/:id", deleteOrder)
	}
}

func createOrder(c *gin.Context) {
	var req model.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.OrderFail(err.Error()))
		return
	}
	data, err := usecase.CreateOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.OrderFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, model.OrderOK("Order berhasil dibuat", data))
}

func getAllOrder(c *gin.Context) {
	status := c.Query("status")
	if status != "" {
		list, err := usecase.GetOrderByStatus(entity.OrderStatus(status))
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.OrderFail(err.Error()))
			return
		}
		c.JSON(http.StatusOK, model.OrderOK("OK", list))
		return
	}
	list, err := usecase.GetAllOrder()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.OrderFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.OrderOK("OK", list))
}

func getOrderByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.OrderFail("ID tidak valid"))
		return
	}
	data, err := usecase.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.OrderFail("Order tidak ditemukan"))
		return
	}
	c.JSON(http.StatusOK, model.OrderOK("OK", data))
}

func getOrderBySPPGID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("sppg_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.OrderFail("ID tidak valid"))
		return
	}
	list, err := usecase.GetOrderBySPPGID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.OrderFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.OrderOK("OK", list))
}

func updateOrderStatus(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.OrderFail("ID tidak valid"))
		return
	}
	var req model.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.OrderFail(err.Error()))
		return
	}
	if err := usecase.UpdateOrderStatus(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.OrderFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.OrderOK("Status order berhasil diupdate", nil))
}

func deleteOrder(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.OrderFail("ID tidak valid"))
		return
	}
	if err := usecase.DeleteOrder(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.OrderFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.OrderOK("Order berhasil dihapus", nil))
}
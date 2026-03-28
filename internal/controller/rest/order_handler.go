package rest

import (
	"fmt"
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
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	data, err := usecase.CreateOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusCreated, model.Created(data))
}

func getAllOrder(c *gin.Context) {
	status := c.Query("status")
	sppgIDStr := c.Query("sppg_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, limit := 1, 10
	fmt.Sscanf(pageStr, "%d", &page)
	fmt.Sscanf(limitStr, "%d", &limit)

	if status == "" && sppgIDStr == "" && startDate == "" && endDate == "" {
		list, err := usecase.GetAllOrder()
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.InternalError())
			return
		}
		c.JSON(http.StatusOK, model.OK(gin.H{"data": list, "total": len(list)}))
		return
	}

	var sppgID *uuid.UUID
	if sppgIDStr != "" {
		id, err := uuid.Parse(sppgIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, model.BadRequest("Invalid sppg_id format"))
			return
		}
		sppgID = &id
	}

	var startPtr, endPtr *string
	if startDate != "" {
		startPtr = &startDate
	}
	if endDate != "" {
		endPtr = &endDate
	}

	list, total, err := usecase.GetOrdersFiltered(status, sppgID, startPtr, endPtr, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}

	c.JSON(http.StatusOK, model.OKPaginated(list, total, page, limit))
}

func getOrderByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	data, err := usecase.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NotFound("Order"))
		return
	}
	c.JSON(http.StatusOK, model.OK(data))
}

func getOrderBySPPGID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("sppg_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	list, err := usecase.GetOrderBySPPGID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OK(list))
}

func updateOrderStatus(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	var req model.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	if err := usecase.UpdateOrderStatus(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OKMessage("Order status updated successfully", nil))
}

func deleteOrder(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}

	// Cek status order dulu
	order, err := usecase.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NotFound("Order"))
		return
	}
	if order.OrderStatus != entity.OrderCancelled && order.OrderStatus != entity.OrderPending {
		c.JSON(http.StatusBadRequest, model.BadRequest("Only pending or cancelled orders can be deleted"))
		return
	}

	if err := usecase.DeleteOrder(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.Deleted())
}
package rest

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func StockRoutes(r *gin.RouterGroup) {
	stock := r.Group("/stocks")
	{
		stock.POST("", createStock)
		stock.GET("", getAllStock)
		stock.GET("/:id", getStockByID)
		stock.GET("/product/:product_id", getStockByProductID)
		stock.GET("/supplier/:supplier_id", getStockBySupplierID)
		stock.PUT("/:id", updateStockQuantity)
		stock.DELETE("/:id", deleteStock)
	}
}

func createStock(c *gin.Context) {
	var req model.CreateStockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.StockFail(err.Error()))
		return
	}
	data, err := usecase.CreateStock(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.StockFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, model.StockOK("Stock berhasil dibuat", data))
}

func getAllStock(c *gin.Context) {
	list, err := usecase.GetAllStock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.StockFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.StockOK("OK", list))
}

func getStockByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.StockFail("ID tidak valid"))
		return
	}
	data, err := usecase.GetStockByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.StockFail("Stock tidak ditemukan"))
		return
	}
	c.JSON(http.StatusOK, model.StockOK("OK", data))
}

func getStockByProductID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.StockFail("ID tidak valid"))
		return
	}
	data, err := usecase.GetStockByProductID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.StockFail("Stock tidak ditemukan"))
		return
	}
	c.JSON(http.StatusOK, model.StockOK("OK", data))
}

func getStockBySupplierID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("supplier_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.StockFail("ID tidak valid"))
		return
	}
	list, err := usecase.GetStockBySupplierID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.StockFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.StockOK("OK", list))
}

func updateStockQuantity(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.StockFail("ID tidak valid"))
		return
	}
	var req model.UpdateStockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.StockFail(err.Error()))
		return
	}
	if err := usecase.UpdateStockQuantity(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.StockFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.StockOK("Stock berhasil diupdate", nil))
}

func deleteStock(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.StockFail("ID tidak valid"))
		return
	}
	if err := usecase.DeleteStock(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.StockFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.StockOK("Stock berhasil dihapus", nil))
}
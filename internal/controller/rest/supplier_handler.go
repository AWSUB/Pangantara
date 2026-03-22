package rest

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SupplierRoutes(r *gin.RouterGroup) {
	supplier := r.Group("/suppliers")
	{
		supplier.POST("", createSupplier)
		supplier.GET("", getAllSupplier)
		supplier.GET("/:id", getSupplierByID)
		supplier.GET("/user/:user_id", getSupplierByUserID)
		supplier.PUT("/:id", updateSupplier)
		supplier.DELETE("/:id", deleteSupplier)
	}
}

func createSupplier(c *gin.Context) {
	var req model.CreateSupplierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.SupplierFail(err.Error()))
		return
	}
	data, err := usecase.CreateSupplier(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.SupplierFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, model.SupplierOK("Supplier berhasil dibuat", data))
}

func getAllSupplier(c *gin.Context) {
	list, err := usecase.GetAllSupplier()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.SupplierFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.SupplierOK("OK", list))
}

func getSupplierByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.SupplierFail("ID tidak valid"))
		return
	}
	data, err := usecase.GetSupplierByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.SupplierFail("Supplier tidak ditemukan"))
		return
	}
	c.JSON(http.StatusOK, model.SupplierOK("OK", data))
}

func getSupplierByUserID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.SupplierFail("ID tidak valid"))
		return
	}
	data, err := usecase.GetSupplierByUserID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.SupplierFail("Supplier tidak ditemukan"))
		return
	}
	c.JSON(http.StatusOK, model.SupplierOK("OK", data))
}

func updateSupplier(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.SupplierFail("ID tidak valid"))
		return
	}
	var req model.UpdateSupplierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.SupplierFail(err.Error()))
		return
	}
	if err := usecase.UpdateSupplier(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.SupplierFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.SupplierOK("Supplier berhasil diupdate", nil))
}

func deleteSupplier(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.SupplierFail("ID tidak valid"))
		return
	}
	if err := usecase.DeleteSupplier(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.SupplierFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.SupplierOK("Supplier berhasil dihapus", nil))
}
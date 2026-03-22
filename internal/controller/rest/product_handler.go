package rest

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProductRoutes(r *gin.RouterGroup) {
	product := r.Group("/products")
	{
		product.POST("", createProduct)
		product.GET("", getAllProduct)
		product.GET("/:id", getProductByID)
		product.GET("/supplier/:supplier_id", getProductBySupplier)
		product.PUT("/:id", updateProduct)
		product.DELETE("/:id", deleteProduct)
	}
}

func createProduct(c *gin.Context) {
	var req model.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ProductFail(err.Error()))
		return
	}
	data, err := usecase.CreateProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ProductFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, model.ProductOK("Product berhasil dibuat", data))
}

func getAllProduct(c *gin.Context) {
	category := c.Query("category")
	if category != "" {
		list, err := usecase.GetProductByCategory(category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ProductFail(err.Error()))
			return
		}
		c.JSON(http.StatusOK, model.ProductOK("OK", list))
		return
	}
	list, err := usecase.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ProductFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.ProductOK("OK", list))
}

func getProductByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ProductFail("ID tidak valid"))
		return
	}
	data, err := usecase.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ProductFail("Product tidak ditemukan"))
		return
	}
	c.JSON(http.StatusOK, model.ProductOK("OK", data))
}

func getProductBySupplier(c *gin.Context) {
	id, err := uuid.Parse(c.Param("supplier_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ProductFail("ID tidak valid"))
		return
	}
	category := c.Query("category")
	if category != "" {
		list, err := usecase.GetProductBySupplierAndCategory(id, category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ProductFail(err.Error()))
			return
		}
		c.JSON(http.StatusOK, model.ProductOK("OK", list))
		return
	}
	list, err := usecase.GetProductBySupplier(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ProductFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.ProductOK("OK", list))
}

func updateProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ProductFail("ID tidak valid"))
		return
	}
	var req model.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ProductFail(err.Error()))
		return
	}
	if err := usecase.UpdateProduct(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.ProductFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.ProductOK("Product berhasil diupdate", nil))
}

func deleteProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ProductFail("ID tidak valid"))
		return
	}
	if err := usecase.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.ProductFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.ProductOK("Product berhasil dihapus", nil))
}
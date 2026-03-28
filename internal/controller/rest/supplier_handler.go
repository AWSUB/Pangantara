package rest

import (
	"net/http"
	"sppg-backend/internal/middleware"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"
	"sppg-backend/pkg/upload"

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
		supplier.DELETE("/:id", middleware.RoleMiddleware("admin"), deleteSupplier)
		supplier.PATCH("/:id/verify", middleware.RoleMiddleware("admin"), verifySupplier)
	}
}

func createSupplier(c *gin.Context) {
	var req model.CreateSupplierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	data, err := usecase.CreateSupplier(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusCreated, model.Created(data))
}

func getAllSupplier(c *gin.Context) {
	keyword := c.Query("keyword")
	category := c.Query("category")
	status := c.Query("status")

	if keyword != "" || category != "" {
		list, err := usecase.SearchSupplier(keyword, category, status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.InternalError())
			return
		}
		c.JSON(http.StatusOK, model.OK(list))
		return
	}

	if status != "" {
		list, err := usecase.GetSupplierByVerificationStatus(status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.InternalError())
			return
		}
		c.JSON(http.StatusOK, model.OK(list))
		return
	}

	list, err := usecase.GetAllSupplier()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OK(list))
}

func getSupplierByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	data, err := usecase.GetSupplierByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NotFound("Supplier"))
		return
	}
	c.JSON(http.StatusOK, model.OK(data))
}

func getSupplierByUserID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	data, err := usecase.GetSupplierByUserID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NotFound("Supplier"))
		return
	}
	c.JSON(http.StatusOK, model.OK(data))
}

func updateSupplier(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	var req model.UpdateSupplierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	if err := usecase.UpdateSupplier(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.Updated())
}

func verifySupplier(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	var req model.VerifySupplierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	if err := usecase.VerifySupplier(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	msg := "Supplier approved successfully"
	if req.Status == "rejected" {
		msg = "Supplier rejected successfully"
	}
	c.JSON(http.StatusOK, model.OKMessage(msg, nil))
}

func deleteSupplier(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	if err := usecase.DeleteSupplier(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.Deleted())
}

func uploadSupplierDoc(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	docType := c.PostForm("document_type")
	if docType != "nib" && docType != "halal" && docType != "other" {
		c.JSON(http.StatusBadRequest, model.BadRequest("document_type must be: nib, halal, or other"))
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("File not found in request"))
		return
	}
	savedPath, err := upload.SaveDocument(file, "supplier")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest(err.Error()))
		return
	}
	if err := usecase.UpdateSupplierDocument(id, docType, savedPath); err != nil {
		_ = upload.DeleteFile(savedPath)
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OKMessage("Document uploaded successfully", gin.H{
		"document_type": docType,
		"file_url":      savedPath,
	}))
}
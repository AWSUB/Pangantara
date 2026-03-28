package rest

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"
	"sppg-backend/pkg/upload"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SupplierDraftRoutes(r *gin.RouterGroup) {
	draft := r.Group("/supplier-draft")
	{
		draft.POST("/save", saveDraft)
		draft.GET("/:user_id", getDraft)
		draft.POST("/submit", submitDraft)
		draft.PATCH("/:user_id/document", uploadDraftDocument)
		draft.DELETE("/:user_id", deleteDraft)
	}
}

func saveDraft(c *gin.Context) {
	var req model.SaveDraftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	data, err := usecase.SaveDraft(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OKMessage("Draft saved successfully", data))
}

func getDraft(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	data, err := usecase.GetDraft(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NotFound("Draft"))
		return
	}
	c.JSON(http.StatusOK, model.OK(data))
}

func submitDraft(c *gin.Context) {
	var req model.SubmitDraftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	data, err := usecase.SubmitDraft(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, model.OKMessage("Registration submitted successfully, waiting for admin verification", data))
}

func uploadDraftDocument(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
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
	savedPath, err := upload.SaveDocument(file, "supplier-draft")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest(err.Error()))
		return
	}
	if err := usecase.UpdateDraftDocument(userID, docType, savedPath); err != nil {
		_ = upload.DeleteFile(savedPath)
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OKMessage("Document uploaded successfully", gin.H{
		"document_type": docType,
		"file_url":      savedPath,
	}))
}

func deleteDraft(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	if err := usecase.DeleteDraft(userID); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.Deleted())
}
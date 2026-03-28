package rest

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SPPGRoutes(r *gin.RouterGroup) {
	sppg := r.Group("/sppg")
	{
		sppg.POST("", createSPPG)
		sppg.GET("", getAllSPPG)
		sppg.GET("/:id", getSPPGByID)
		sppg.GET("/user/:user_id", getSPPGByUserID)
		sppg.PUT("/:id", updateSPPG)
		sppg.DELETE("/:id", deleteSPPG)
	}
}

func createSPPG(c *gin.Context) {
	var req model.CreateSPPGRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	data, err := usecase.CreateSPPG(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusCreated, model.Created(data))
}

func getAllSPPG(c *gin.Context) {
	list, err := usecase.GetAllSPPG()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OK(list))
}

func getSPPGByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	data, err := usecase.GetSPPGByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NotFound("SPPG"))
		return
	}
	c.JSON(http.StatusOK, model.OK(data))
}

func getSPPGByUserID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	list, err := usecase.GetSPPGByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OK(list))
}

func updateSPPG(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	var req model.UpdateSPPGRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	if err := usecase.UpdateSPPG(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.Updated())
}

func deleteSPPG(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest("Invalid ID format"))
		return
	}
	if err := usecase.DeleteSPPG(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.Deleted())
}
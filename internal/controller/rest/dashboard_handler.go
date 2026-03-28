package rest

import (
	"net/http"
	"sppg-backend/internal/middleware"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func DashboardRoutes(r *gin.RouterGroup) {
	dashboard := r.Group("/dashboard")
	dashboard.Use(middleware.RoleMiddleware("admin"))
	{
		dashboard.GET("/summary", getDashboardSummary)
	}
}

func getDashboardSummary(c *gin.Context) {
	summary, err := usecase.GetDashboardSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.InternalError())
		return
	}
	c.JSON(http.StatusOK, model.OK(summary))
}
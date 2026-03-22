package rest

import (
	"sppg-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	// Public routes (tidak perlu token)
	AuthRoutes(api)

	// Protected routes (perlu token)
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		UserRoutes(protected)
		SPPGRoutes(protected)
		SupplierRoutes(protected)
		ProductRoutes(protected)
		StockRoutes(protected)
		OrderRoutes(protected)
		TransactionRoutes(protected)
	}
}
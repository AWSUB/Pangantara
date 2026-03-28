package rest

import (
	"net/http"
	"sppg-backend/internal/middleware"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.RouterGroup) {
	payment := r.Group("/payment")
	{
		payment.POST("/create", middleware.RoleMiddleware("sppg", "admin"), createPayment)
	}
}

func WebhookRoutes(r *gin.RouterGroup) {
	webhook := r.Group("/webhook")
	{
		webhook.POST("/midtrans", midtransWebhook)
	}
}

func createPayment(c *gin.Context) {
	var req model.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	resp, err := usecase.CreatePayment(req.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.BadRequest(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.OKMessage("Payment created successfully", gin.H{
		"token":        resp.Token,
		"redirect_url": resp.RedirectURL,
	}))
}

func midtransWebhook(c *gin.Context) {
	var notif model.MidtransNotification
	if err := c.ShouldBindJSON(&notif); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationError(err.Error()))
		return
	}
	if err := usecase.HandleMidtransNotification(notif); err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.OKMessage("Notification processed successfully", nil))
}
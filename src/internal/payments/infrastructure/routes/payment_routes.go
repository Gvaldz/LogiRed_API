package routes

import (
	"logired/src/internal/payments/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type PaymentRoutes struct {
	controller     *controllers.PaymentController
	authMiddleware gin.HandlerFunc
}

func NewPaymentRoutes(controller *controllers.PaymentController, authMiddleware gin.HandlerFunc) *PaymentRoutes {
	return &PaymentRoutes{controller: controller, authMiddleware: authMiddleware}
}

func (r *PaymentRoutes) AttachRoutes(router *gin.Engine) {
	group := router.Group("/payments")
	group.Use(r.authMiddleware)
	group.POST("/create-intent", r.controller.CreateIntent)

	// Webhook sin auth (Stripe lo llama directamente)
	router.POST("/payments/webhook", r.controller.WebhookHandler)
}

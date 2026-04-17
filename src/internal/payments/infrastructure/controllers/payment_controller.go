package controllers

import (
	"logired/src/internal/payments/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	createPaymentIntent *application.CreatePaymentIntent
}

func NewPaymentController(createPaymentIntent *application.CreatePaymentIntent) *PaymentController {
	return &PaymentController{createPaymentIntent: createPaymentIntent}
}

func (ctrl *PaymentController) CreateIntent(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	clientID := userIDInterface.(int32)

	var req struct {
		RideId int32   `json:"ride_id"`
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.createPaymentIntent.Execute(req.RideId, clientID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"client_secret":     result.ClientSecret,
		"payment_intent_id": result.PaymentIntentId,
	})
}

func (ctrl *PaymentController) WebhookHandler(c *gin.Context) {
	// Por ahora solo responde OK — lo expandimos después
	c.JSON(http.StatusOK, gin.H{"received": true})
}

func init() {
	_ = strconv.Itoa // evita import no usado
}

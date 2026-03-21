package controllers

import (
	"logired/src/internal/reviews/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetReviewsByDriverController struct {
	getReviewsByDriver *application.GetReviewsByDriver
}

func NewGetReviewsByDriverController(get *application.GetReviewsByDriver) *GetReviewsByDriverController {
	return &GetReviewsByDriverController{getReviewsByDriver: get}
}

func (ctrl *GetReviewsByDriverController) GetByDriver(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	driverID := userIDInterface.(int32)

	reviews, err := ctrl.getReviewsByDriver.Execute(driverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}
package controllers

import (
	"logired/src/internal/reviews/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetReviewsByPassangerController struct {
	getReviewsByPassanger *application.GetReviewsByPassanger
}

func NewGetReviewsByPassangerController(get *application.GetReviewsByPassanger) *GetReviewsByPassangerController {
	return &GetReviewsByPassangerController{getReviewsByPassanger: get}
}

func (ctrl *GetReviewsByPassangerController) GetByPassanger(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	passangerID := userIDInterface.(int32)

	reviews, err := ctrl.getReviewsByPassanger.Execute(passangerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}
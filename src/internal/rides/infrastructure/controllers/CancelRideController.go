package controllers

import (
	"logired/src/internal/rides/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CancelRideController struct {
	cancelRide *application.CancelRide
}

func NewCancelRideController(cancel *application.CancelRide) *CancelRideController {
	return &CancelRideController{cancelRide: cancel}
}

func (ctrl *CancelRideController) Cancel(c *gin.Context) {
	idParam := c.Param("id")
	idRide, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	clientID := userIDInterface.(int32)

	if err := ctrl.cancelRide.Execute(int32(idRide), clientID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Viaje cancelado correctamente"})
}
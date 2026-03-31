package controllers

import (
	"logired/src/internal/devices/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceTokenController struct {
	saveToken *application.SaveDeviceToken
}

func NewDeviceTokenController(uc *application.SaveDeviceToken) *DeviceTokenController {
	return &DeviceTokenController{saveToken: uc}
}

func (ctrl *DeviceTokenController) Register(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	var body struct {
		Token      string `json:"fcm_token"   binding:"required"`
		DeviceName string `json:"device_name"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.saveToken.Execute(userID.(int32), body.Token, body.DeviceName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dispositivo registrado correctamente"})
}
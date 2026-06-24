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

// Register godoc
// @Summary      Registrar dispositivo para notificaciones
// @Description  Asocia un token FCM y nombre de dispositivo al usuario autenticado
// @Tags         devices
// @Accept       json
// @Produce      json
// @Param        body body map[string]interface{} true "Datos del dispositivo" Example({"fcm_token":"abc123","device_name":"iPhone"})
// @Security     ApiKeyAuth
// @Success      200 {object} map[string]string "mensaje de éxito"
// @Failure      400 {object} map[string]string "error en la solicitud"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /devices/token [put]
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



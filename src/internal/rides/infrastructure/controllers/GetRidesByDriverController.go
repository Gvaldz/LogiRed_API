package controllers

import (
	"logired/src/internal/rides/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetRidesByDriverController struct {
	getRidesByDriver *application.GetRidesByDriver
}

func NewGetRidesByDriverController(get *application.GetRidesByDriver) *GetRidesByDriverController {
	return &GetRidesByDriverController{getRidesByDriver: get}
}

// GetByDriver godoc
// @Summary      Obtener viajes asignados al conductor autenticado
// @Description  Devuelve todos los viajes que el conductor ha aceptado
// @Tags         rides
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200 {object} map[string]interface{} "lista de viajes"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /rides/driver/me [get]
func (ctrl *GetRidesByDriverController) GetByDriver(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	driverID := userIDInterface.(int32)

	rides, err := ctrl.getRidesByDriver.Execute(driverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rides": rides})
}
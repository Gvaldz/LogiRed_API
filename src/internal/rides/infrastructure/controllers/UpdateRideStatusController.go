package controllers

import (
	"logired/src/internal/rides/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateRideStatusController struct {
	UpdateRideStatus *application.UpdateRideStatus
}

func NewUpdateRideStatusController(update *application.UpdateRideStatus) *UpdateRideStatusController {
	return &UpdateRideStatusController{UpdateRideStatus: update}
}

// Update godoc
// @Summary      Actualizar estado de un viaje
// @Description  Cambia el estado de un viaje (solo el cliente o conductor autorizado)
// @Tags         rides
// @Accept       json
// @Produce      json
// @Param        id path int true "ID del viaje"
// @Param        body body map[string]interface{} true "Nuevo estado" Example({"status":2})
// @Security     ApiKeyAuth
// @Success      200 {object} map[string]string "mensaje de éxito"
// @Failure      400 {object} map[string]string "error en la solicitud"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      404 {object} map[string]string "viaje no encontrado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /rides/{id}/status [put]
func (ctrl *UpdateRideStatusController) Update(c *gin.Context) {
	idParam := c.Param("id")
	rideID, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de viaje inválido"})
		return
	}

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	userID := userIDInterface.(int32)
	_ = userID
	var req struct {
		Status int32 `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.UpdateRideStatus.Execute(int32(rideID), req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 5. Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Estado del viaje actualizado correctamente"})
}

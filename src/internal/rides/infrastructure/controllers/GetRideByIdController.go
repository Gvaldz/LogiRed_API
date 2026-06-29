package controllers

import (
	"logired/src/internal/rides/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetRideByIdController struct {
	getRideById *application.GetRideById
}

func NewGetRideByIdController(get *application.GetRideById) *GetRideByIdController {
	return &GetRideByIdController{getRideById: get}
}

// GetById godoc
// @Summary      Obtener un viaje por ID
// @Description  Devuelve los detalles de un viaje específico
// @Tags         rides
// @Security     Bearer
// @Produce      json
// @Param        id path int true "ID del viaje"
// @Success      200 {object} map[string]interface{} "datos del viaje"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      404 {object} map[string]string "viaje no encontrado"
// @Router       /rides/{id} [get]
func (ctrl *GetRideByIdController) GetById(c *gin.Context) {
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
	userID := userIDInterface.(int32)

	ride, err := ctrl.getRideById.Execute(int32(idRide))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userTypeInterface, _ := c.Get("userType")
	userType, _ := userTypeInterface.(int)

	if userType == 1 {
		// Clientes solo pueden ver sus propios viajes
		if ride.IdClient != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para ver este viaje"})
			return
		}
	} else if userType == 2 {
		// Conductores pueden ver viajes disponibles o asignados a ellos
		isAssigned := ride.IdDriver != nil && *ride.IdDriver == userID
		isAvailable := ride.IdStatus == 6
		if !isAssigned && !isAvailable {
			c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para ver este viaje"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"ride": ride})
}
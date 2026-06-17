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
// @Produce      json
// @Param        id path int true "ID del viaje"
// @Success      200 {object} map[string]interface{} "datos del viaje"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      404 {object} map[string]string "viaje no encontrado"
// @Router       /rides/{id} [get]
func (ctrl *GetRideByIdController) GetById(c *gin.Context) {
	idParam := c.Param("id")
	idRide, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ride, err := ctrl.getRideById.Execute(int32(idRide))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ride": ride})
}
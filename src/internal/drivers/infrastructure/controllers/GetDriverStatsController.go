package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"logired/src/internal/drivers/application"
)

type GetDriverStatsController struct {
	getDriverStats *application.GetDriverStats
}

func NewGetDriverStatsController(getDriverStats *application.GetDriverStats) *GetDriverStatsController {
	return &GetDriverStatsController{getDriverStats: getDriverStats}
}

// GetStats godoc
// @Summary      Obtener estadísticas del conductor
// @Description  Devuelve estadísticas detalladas del conductor para machine learning (rating, viajes completados, cancelados, etc.)
// @Tags         drivers
// @Security     Bearer
// @Produce      json
// @Param        id path int true "ID del conductor"
// @Success      200 {object} map[string]interface{} "estadísticas del conductor"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      404 {object} map[string]string "conductor no encontrado"
// @Router       /drivers/{id}/stats [get]
func (ctrl *GetDriverStatsController) GetStats(c *gin.Context) {
	idStr := c.Param("id")
	driverID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de conductor inválido"})
		return
	}

	stats, err := ctrl.getDriverStats.Execute(int32(driverID))
	if err != nil {
		if err.Error() == "conductor no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stats": stats})
}

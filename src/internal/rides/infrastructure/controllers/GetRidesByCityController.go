package controllers

import (
    "logired/src/internal/rides/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type GetRideByCityController struct {
    getRideByCity *application.GetRidesByCity
}

func NewGetRideByCityController(get *application.GetRidesByCity) *GetRideByCityController {
    return &GetRideByCityController{getRideByCity: get}
}

// GetAvailable godoc
// @Summary      Obtener todos los viajes disponibles
// @Description  Devuelve todos los viajes activos (estado 6) disponibles para cualquier conductor. Requiere autenticación.
// @Tags         rides
// @Security     Bearer
// @Produce      json
// @Success      200 {object} map[string]interface{} "lista de viajes disponibles"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /rides/available [get]
func (ctrl *GetRideByCityController) GetByCity(c *gin.Context) {
    rides, err := ctrl.getRideByCity.Execute("")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"rides": rides})
}
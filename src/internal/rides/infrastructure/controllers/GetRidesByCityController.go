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

// GetByCity godoc
// @Summary      Obtener viajes disponibles por ciudad
// @Description  Devuelve viajes activos (estado 6) en la ciudad especificada
// @Tags         rides
// @Produce      json
// @Param        city path string true "Nombre de la ciudad"
// @Success      200 {object} map[string]interface{} "lista de viajes"
// @Failure      400 {object} map[string]string "ciudad no especificada"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /rides/city/{city} [get]
func (ctrl *GetRideByCityController) GetByCity(c *gin.Context) {
    city := c.Param("city")
    if city == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ciudad no especificada"})
        return
    }

    rides, err := ctrl.getRideByCity.Execute(city)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"rides": rides})
}
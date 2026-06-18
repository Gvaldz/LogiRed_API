package controllers


/*
import (
    "logired/src/internal/drivers/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type GetDriversByCityController struct {
    getDriversByCity *application.GetDriversByCity
}

func NewGetRideByCityController(get *application.GetDriversByCity) *GetDriversByCityController {
    return &GetDriversByCityController{getDriversByCity: get}
}

// GetByCity godoc
// @Summary      Obtener conductores por ciudad
// @Description  Devuelve una lista de conductores que trabajan en una ciudad (búsqueda por LIKE)
// @Tags         drivers
// @Produce      json
// @Param        city path string true "Nombre de la ciudad (p.ej. 'Tuxtla')"
// @Success      200 {object} map[string]interface{} "lista de conductores"
// @Failure      400 {object} map[string]string "ciudad no especificada"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /drivers/city/{city} [get]
func (ctrl *GetDriversByCityController) GetByCity(c *gin.Context) {
    city := c.Param("city")
    if city == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ciudad no especificada"})
        return
    }

    drivers, err := ctrl.getDriversByCity.Execute(city)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"drivers": drivers})
}
    */
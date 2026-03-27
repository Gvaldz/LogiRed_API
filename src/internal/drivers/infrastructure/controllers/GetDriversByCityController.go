package controllers

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
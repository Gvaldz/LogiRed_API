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
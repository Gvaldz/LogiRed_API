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

func (ctrl *GetRideByIdController) GetById(c *gin.Context) {
	idParam := c.Param("id")
	idRide, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Se asume que cualquier usuario autenticado puede ver un viaje por ID
	ride, err := ctrl.getRideById.Execute(int32(idRide))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ride": ride})
}
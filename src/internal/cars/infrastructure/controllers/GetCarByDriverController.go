package controllers

import (
	"logired/src/internal/cars/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetCarsByDriverController struct {
	getCarsByDriver *application.GetCarsByDriver
}

func NewGetCarsByDriverController(get *application.GetCarsByDriver) *GetCarsByDriverController {
	return &GetCarsByDriverController{
		getCarsByDriver: get,
	}
}

func (ctrl *GetCarsByDriverController) GetByDriver(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	idDriver := userIDInterface.(int32)

	cars, err := ctrl.getCarsByDriver.Execute(idDriver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cars": cars})
}
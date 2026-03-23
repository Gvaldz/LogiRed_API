package controllers

import (
	"logired/src/internal/drivers/application"
	"logired/src/internal/drivers/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateDriverController struct {
	updateDriver *application.UpdateDriver
}

func NewUpdateDriverController(update *application.UpdateDriver) *UpdateDriverController {
	return &UpdateDriverController{updateDriver: update}
}

func (ctrl *UpdateDriverController) Update(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	userID := userIDInterface.(int32)

	var req struct {
		Rating 		float32 `json:"rating"`
		Citywork  	string  `json:"citywork"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driver := entities.Driver{
		IdUser: userID,
		Rating: req.Rating,
		Citywork:  req.Citywork,
	}

	if err := ctrl.updateDriver.Execute(driver); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Perfil de conductor actualizado correctamente",
		"driver":  driver,
	})
}
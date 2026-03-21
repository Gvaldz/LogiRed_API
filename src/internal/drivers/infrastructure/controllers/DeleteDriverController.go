package controllers

import (
	"logired/src/internal/drivers/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteDriverController struct {
	deleteDriver *application.DeleteDriver
}

func NewDeleteDriverController(delete *application.DeleteDriver) *DeleteDriverController {
	return &DeleteDriverController{deleteDriver: delete}
}

func (ctrl *DeleteDriverController) Delete(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	userID := userIDInterface.(int32)

	if err := ctrl.deleteDriver.Execute(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perfil de conductor eliminado correctamente"})
}
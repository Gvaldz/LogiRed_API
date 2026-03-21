package controllers

import (
	"logired/src/internal/cars/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCarController struct {
	deleteCar *application.DeleteCar
}

func NewDeleteCarController(delete *application.DeleteCar) *DeleteCarController {
	return &DeleteCarController{
		deleteCar: delete,
	}
}

func (ctrl *DeleteCarController) Delete(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	idDriver := userIDInterface.(int32)

	idStr := c.Param("id")
	idCar, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.deleteCar.Execute(int32(idCar), idDriver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Carro eliminado correctamente"})
}
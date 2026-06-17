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

// Delete godoc
// @Summary      Eliminar un vehículo
// @Description  Solo el conductor propietario puede eliminar su carro
// @Tags         cars
// @Produce      json
// @Param        id path int true "ID del carro"
// @Security     ApiKeyAuth
// @Success      200 {object} map[string]string "mensaje de éxito"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      404 {object} map[string]string "carro no encontrado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /cars/{id} [delete]
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
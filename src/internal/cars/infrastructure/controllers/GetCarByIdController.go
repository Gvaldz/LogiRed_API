package controllers

import (
	"logired/src/internal/cars/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetCarByIdController struct {
	getCarById *application.GetCarById
}

func NewGetCarByIdController(get *application.GetCarById) *GetCarByIdController {
	return &GetCarByIdController{
		getCarById: get,
	}
}

// GetById godoc
// @Summary      Obtener un vehículo por ID
// @Tags         cars
// @Produce      json
// @Param        id path int true "ID del carro"
// @Success      200 {object} map[string]interface{} "datos del carro"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      404 {object} map[string]string "carro no encontrado"
// @Router       /cars/{id} [get]
func (ctrl *GetCarByIdController) GetById(c *gin.Context) {


	idStr := c.Param("id")
	idCar, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	car, err := ctrl.getCarById.Execute(int32(idCar))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"car": car})
}
package controllers

import (
	"logired/src/internal/cars/application"
	"logired/src/internal/cars/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateCarController struct {
	updateCar *application.UpdateCar
}

func NewUpdateCarController(update *application.UpdateCar) *UpdateCarController {
	return &UpdateCarController{
		updateCar: update,
	}
}

func (ctrl *UpdateCarController) Update(c *gin.Context) {
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

	var request struct {
		CarRegistration string `json:"car_registration"`
		Brand           string `json:"brand"`
		Model           string `json:"model"`
		Color           string `json:"color"`
		MaxCapacity     int32  `json:"max_capacity"`
		Image           string `json:"image"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := entities.Car{
		IdCar:           int32(idCar),
		IdDriver:        idDriver,
		CarRegistration: request.CarRegistration,
		Brand:           request.Brand,
		Model:           request.Model,
		Color:           request.Color,
		MaxCapacity:     request.MaxCapacity,
		Image:           request.Image,
	}

	err = ctrl.updateCar.Execute(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Carro actualizado correctamente", "car": car})
}
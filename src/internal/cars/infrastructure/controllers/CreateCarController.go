package controllers

import (
	"logired/src/internal/cars/application"
	"logired/src/internal/cars/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCarController struct {
	createCar *application.CreateCar
}

func NewCreateCarController(create *application.CreateCar) *CreateCarController {
	return &CreateCarController{
		createCar: create,
	}
}

func (ctrl *CreateCarController) Create(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	idDriver := userIDInterface.(int32)

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
		IdDriver:        idDriver,
		CarRegistration: request.CarRegistration,
		Brand:           request.Brand,
		Model:           request.Model,
		Color:           request.Color,
		MaxCapacity:     request.MaxCapacity,
		Image:           request.Image,
	}

	err := ctrl.createCar.Execute(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Carro creado correctamente", "car": car})
}
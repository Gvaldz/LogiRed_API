package controllers

import (
	"logired/src/internal/drivers/application"
	"logired/src/internal/drivers/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateDriverController struct {
	createDriver *application.CreateDriver
}

func NewCreateDriverController(create *application.CreateDriver) *CreateDriverController {
	return &CreateDriverController{createDriver: create}
}

func (ctrl *CreateDriverController) Create(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	userID := userIDInterface.(int32)

	var req struct {
		Rating 		float32 `json:"rating"`
		Image  		string  `json:"image"`
		Citywork 	string `json:"citywork"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driver := entities.Driver{
		IdUser: userID,
		Rating: req.Rating,
		Citywork: req.Citywork,
	}

	if err := ctrl.createDriver.Execute(driver); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Perfil de conductor creado correctamente",
		"driver":  driver,
	})
}
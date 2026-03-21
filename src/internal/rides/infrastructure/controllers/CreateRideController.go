package controllers

import (
	"logired/src/internal/rides/application"
	"logired/src/internal/rides/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRideController struct {
	createRide *application.CreateRide
}

func NewCreateRideController(create *application.CreateRide) *CreateRideController {
	return &CreateRideController{createRide: create}
}

func (ctrl *CreateRideController) Create(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	clientID := userIDInterface.(int32)

	var req struct {
		Date        string `json:"date"`
		Hour        string `json:"hour"`
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ride := entities.Ride{
		IdClient:    clientID,
		Date:        req.Date,
		Hour:        req.Hour,
		Origin:      req.Origin,
		Destination: req.Destination,
		Description: req.Description,
	}

	if err := ctrl.createRide.Execute(ride); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Viaje creado correctamente", "ride": ride})
}
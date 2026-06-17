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

// Create godoc
// @Summary      Crear un nuevo viaje
// @Description  Un cliente crea un viaje para transporte de paquetes
// @Tags         rides
// @Accept       json
// @Produce      json
// @Param        body body map[string]interface{} true "Datos del viaje" Example({"origin_city":"Tuxtla","origin_address":"Calle 1","origin_lat":12.34,"origin_lng":-56.78,"destination_address":"Calle 2","destination_lat":12.35,"destination_lng":-56.79,"distance_km":10.5,"date":"2025-05-20","hour":"14:30","approx_weight":15.0,"description":"Paquete frágil"})
// @Security     ApiKeyAuth
// @Success      201 {object} map[string]interface{} "mensaje y viaje creado"
// @Failure      400 {object} map[string]string "error en los datos"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /rides [post]
func (ctrl *CreateRideController) Create(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	clientID := userIDInterface.(int32)

	var req struct {
		OriginCity         string  `json:"origin_city"`
		OriginAddress      string  `json:"origin_address"`
		OriginLat          float64 `json:"origin_lat"`
		OriginLng          float64 `json:"origin_lng"`
		DestinationAddress string  `json:"destination_address"`
		DestinationLat     float64 `json:"destination_lat"`
		DestinationLng     float64 `json:"destination_lng"`
		DistanceKm         float64 `json:"distance_km"`
		Date               string  `json:"date"`
		Hour               string  `json:"hour"`
		AproxWeight        float64 `json:"approx_weight"`
		Description        string  `json:"description"`
		IdStatus           int32   `json:"idstatus"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ride := entities.Ride{
		IdClient:          clientID,
		OriginCity:        req.OriginCity,
		OriginAddress:     req.OriginAddress,
		OriginLat:         req.OriginLat,
		OriginLng:         req.OriginLng,
		DestinationAddres: req.DestinationAddress,
		DestinationLat:    req.DestinationLat,
		DestinationLng:    req.DestinationLng,
		DistanceKm:        req.DistanceKm,
		Date:              req.Date,
		Hour:              req.Hour,
		AproxWeight:       req.AproxWeight,
		Description:       req.Description,
		IdStatus:          6,
	}

	if err := ctrl.createRide.Execute(ride); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Viaje creado correctamente", "ride": ride})
}

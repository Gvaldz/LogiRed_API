package controllers

import (
    "net/http"
    "logired/src/internal/rides/application"
    "logired/src/internal/rides/domain/entities"
    "logired/src/internal/rides/infrastructure/dto"
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
// @Param        body body dto.CreateRideRequest true "Datos del viaje"
// @Security     Bearer
// @Success      201 {object} dto.CreateRideResponse
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /rides [post]
func (ctrl *CreateRideController) Create(c *gin.Context) {
    userIDInterface, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }

    userTypeInterface, _ := c.Get("userType")
    if userType, ok := userTypeInterface.(int); ok && userType == 2 {
        c.JSON(http.StatusForbidden, gin.H{"error": "Los conductores no pueden publicar viajes"})
        return
    }

    clientID := userIDInterface.(int32)

    var req dto.CreateRideRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ride := entities.Ride{
        IdClient:          clientID,
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

    // Convertir entidad a DTO de respuesta
    resp := dto.RideResponse{
        IdRide:             ride.IdRide,
        IdClient:           ride.IdClient,
        IdDriver:           ride.IdDriver,
        OriginAddress:      ride.OriginAddress,
        OriginLat:          ride.OriginLat,
        OriginLng:          ride.OriginLng,
        DestinationAddress: ride.DestinationAddres,
        DestinationLat:     ride.DestinationLat,
        DestinationLng:     ride.DestinationLng,
        DistanceKm:         ride.DistanceKm,
        Date:               ride.Date,
        Hour:               ride.Hour,
        AproxWeight:        ride.AproxWeight,
        Description:        ride.Description,
        IdStatus:           ride.IdStatus,
    }

    c.JSON(http.StatusCreated, dto.CreateRideResponse{
        Message: "Viaje creado correctamente",
        Ride:    resp,
    })
}
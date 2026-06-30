package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	tokenService "logired/src/core/services/auth/domain"
	rideDomain "logired/src/internal/rides/domain"
	"logired/src/internal/tracking/infrastructure"
)

const RideStatusInProcess int32 = 3

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}

type TrackingController struct {
	hub          *infrastructure.Hub
	tokenService tokenService.TokenService
	rideRepo     rideDomain.IRide
}

func NewTrackingController(
	hub *infrastructure.Hub,
	tokenService tokenService.TokenService,
	rideRepo rideDomain.IRide,
) *TrackingController {
	return &TrackingController{
		hub:          hub,
		tokenService: tokenService,
		rideRepo:     rideRepo,
	}
}

func (ctrl *TrackingController) authenticate(c *gin.Context) (int32, bool) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token requerido"})
		return 0, false
	}
	userID, err := ctrl.tokenService.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return 0, false
	}
	return userID, true
}

// parseRideID extrae y valida el ID del viaje de la URL
func (ctrl *TrackingController) parseRideID(c *gin.Context) (int32, bool) {
	idParam := c.Param("id")
	rideID, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de viaje inválido"})
		return 0, false
	}
	return int32(rideID), true
}
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// PublishLocation godoc
// @Summary      Conectar WebSocket para publicar ubicación
// @Description  Endpoint para que el conductor inicie una conexión WebSocket y transmita su ubicación en un viaje en proceso. (Nota: Swagger UI no puede probar WebSockets interactivos).
// @Tags         tracking
// @Param        id path int true "ID del viaje"
// @Param        token query string true "Token JWT de autenticación"
// @Success      101 {string} string "Switching Protocols (Conexión WebSocket establecida)"
// @Failure      400 {object} map[string]string "ID de viaje inválido"
// @Failure      401 {object} map[string]string "Token requerido o inválido"
// @Failure      403 {object} map[string]string "No autorizado (viaje no en proceso o no eres el conductor)"
// @Failure      404 {object} map[string]string "Viaje no encontrado"
// @Router       /ws/rides/{id}/publish [get]
func (ctrl *TrackingController) PublishLocation(c *gin.Context) {
	userID, ok := ctrl.authenticate(c)
	if !ok {
		return
	}
	rideID, ok := ctrl.parseRideID(c)
	if !ok {
		return
	}

	ride, err := ctrl.rideRepo.GetRideById(rideID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "viaje no encontrado"})
		return
	}

	if ride.IdStatus != RideStatusInProcess {
		c.JSON(http.StatusForbidden, gin.H{"error": "el viaje no está en proceso"})
		return
	}

	if ride.IdDriver == nil || *ride.IdDriver != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "no eres el conductor de este viaje"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("[WS] error al upgradear: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("[WS] conductor %d publicando en viaje %d", userID, rideID)

	const readTimeout = 60 * time.Second

	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(readTimeout))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(readTimeout))
		return nil
	})

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("[WS] conductor %d desconectado: %v", userID, err)
			break
		}
		conn.SetReadDeadline(time.Now().Add(readTimeout))
		ctrl.hub.Publish(rideID, message)
	}

	ctrl.hub.Publish(rideID, []byte(`{"event":"driver_disconnected"}`))
	ctrl.hub.CloseAll(rideID)
}
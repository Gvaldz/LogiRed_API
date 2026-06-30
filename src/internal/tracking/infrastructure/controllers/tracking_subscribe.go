package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"logired/src/internal/tracking/infrastructure"
)

// SubscribeLocation godoc
// @Summary      Conectar WebSocket para recibir ubicación
// @Description  Endpoint para que el cliente inicie una conexión WebSocket y reciba la ubicación del conductor en tiempo real. (Nota: Swagger UI no puede probar WebSockets interactivos).
// @Tags         tracking
// @Param        id path int true "ID del viaje"
// @Param        token query string true "Token JWT de autenticación"
// @Success      101 {string} string "Switching Protocols (Conexión WebSocket establecida)"
// @Failure      400 {object} map[string]string "ID de viaje inválido"
// @Failure      401 {object} map[string]string "Token requerido o inválido"
// @Failure      403 {object} map[string]string "No autorizado (viaje no en proceso o no eres el cliente)"
// @Failure      404 {object} map[string]string "Viaje no encontrado"
// @Router       /ws/rides/{id}/subscribe [get]
func (ctrl *TrackingController) SubscribeLocation(c *gin.Context) {
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
	if ride.IdClient != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "no eres el cliente de este viaje"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("[WS] error al upgradear: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("[WS] cliente %d suscrito a viaje %d", userID, rideID)

	sub := &infrastructure.Subscriber{Send: make(chan []byte, 16)}
	ctrl.hub.Subscribe(rideID, sub)
	defer ctrl.hub.Unsubscribe(rideID, sub)

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-sub.Send:
			if !ok {
				return
			}
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		case <-ticker.C:
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case <-done:
			log.Printf("[WS] cliente %d cerró conexión", userID)
			return
		}
	}
}
package controllers

import (
    "log"
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"

    rideDomain "logired/src/internal/rides/domain"
    tokenService "logired/src/internal/services/auth/domain"
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

func (ctrl *TrackingController) parseRideID(c *gin.Context) (int32, bool) {
    idParam := c.Param("id")
    rideID, err := strconv.ParseInt(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de viaje inválido"})
        return 0, false
    }
    return int32(rideID), true
}

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

    // AJUSTA estos nombres si tu entidad usa otros
    if ride.IdStatus != RideStatusInProcess {
        c.JSON(http.StatusForbidden, gin.H{"error": "el viaje no está en proceso"})
        return
    }
    if ride.IdDriver != userID {
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

    conn.SetReadLimit(512)
    conn.SetReadDeadline(time.Now().Add(60 * time.Second))
    conn.SetPongHandler(func(string) error {
        conn.SetReadDeadline(time.Now().Add(60 * time.Second))
        return nil
    })

    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            log.Printf("[WS] conductor %d desconectado: %v", userID, err)
            break
        }
        ctrl.hub.Publish(rideID, message)
    }
}

// SubscribeLocation - usado por el cliente para recibir la ubicación del conductor
// GET /ws/rides/:id/subscribe?token=...
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

    // Detecta si el cliente cierra la conexión
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
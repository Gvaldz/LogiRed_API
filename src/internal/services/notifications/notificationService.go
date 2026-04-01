package notifications

import (
	"context"
	"fmt"
	"log"
	"logired/src/core"
	"logired/src/internal/devices/domain"
)

type NotificationService struct {
	fcm        *core.FCMService
	deviceRepo domain.DeviceRepository
}

func NewNotificationService(fcm *core.FCMService, deviceRepo domain.DeviceRepository) *NotificationService {
	return &NotificationService{fcm: fcm, deviceRepo: deviceRepo}
}

func (s *NotificationService) sendAsync(userID int32, title, body string, data map[string]string) {
	go func() {
		tokens, err := s.deviceRepo.GetTokensByUser(userID)
		if err != nil || len(tokens) == 0 {
			log.Printf("[FCM] usuario %d sin dispositivos: %v", userID, err)
			return
		}

		errs := s.fcm.SendToMany(context.Background(), tokens, core.FCMMessage{
			Title: title,
			Body:  body,
			Data:  data,
		})

		for _, err := range errs {
			if core.IsInvalidToken(err) {
				log.Printf("[FCM] token inválido: %v", err)
			}
		}
	}()
}

// ── Viajes ────────────────────────────────────────────────────────

func (s *NotificationService) NuevoViaje(driverID int32, rideID int32) {
	s.sendAsync(driverID,
		"Nuevo viaje disponible 🚗",
		"Hay un viaje en tu ciudad, ¡acepta antes que alguien más!",
		map[string]string{"type": "new_ride", "ride_id": fmt.Sprintf("%d", rideID)},
	)
}

func (s *NotificationService) ViajeAgendado(clientID int32, rideID int32) {
	s.sendAsync(clientID,
		"¡Viaje agendado! 📅",
		"Un conductor aceptó tu viaje",
		map[string]string{"type": "ride_scheduled", "ride_id": fmt.Sprintf("%d", rideID)},
	)
}

func (s *NotificationService) ViajeEnCamino(clientID int32, rideID int32) {
	s.sendAsync(clientID,
		"El conductor está en camino 🛣️",
		"Tu conductor ya va hacia el punto de origen",
		map[string]string{"type": "ride_on_the_way", "ride_id": fmt.Sprintf("%d", rideID)},
	)
}

func (s *NotificationService) ViajeIniciado(clientID int32, rideID int32) {
	s.sendAsync(clientID,
		"¡Tu viaje ha iniciado! 🚀",
		"El viaje está en proceso",
		map[string]string{"type": "ride_started", "ride_id": fmt.Sprintf("%d", rideID)},
	)
}

func (s *NotificationService) ViajeCancelado(userID int32, rideID int32) {
	s.sendAsync(userID,
		"Viaje cancelado ❌",
		"El viaje ha sido cancelado",
		map[string]string{"type": "ride_cancelled", "ride_id": fmt.Sprintf("%d", rideID)},
	)
}

func (s *NotificationService) ViajeFinalizado(clientID int32, rideID int32) {
	s.sendAsync(clientID,
		"¡Viaje finalizado! 🏁",
		"¿Cómo estuvo? No olvides calificar al conductor",
		map[string]string{"type": "ride_finished", "ride_id": fmt.Sprintf("%d", rideID)},
	)
}

// ── Propuestas ────────────────────────────────────────────────────

func (s *NotificationService) NuevaPropuesta(clientID int32, rideID int32) {
	s.sendAsync(clientID,
		"Nueva propuesta recibida 📋",
		"Un conductor hizo una propuesta en tu viaje",
		map[string]string{"type": "new_proposal", "ride_id": fmt.Sprintf("%d", rideID)},
	)
}

func (s *NotificationService) PropuestaAceptada(driverID int32, rideID int32) {
	s.sendAsync(driverID,
		"¡Propuesta aceptada! ✅",
		"El cliente aceptó tu propuesta, prepárate para el viaje",
		map[string]string{"type": "proposal_accepted", "ride_id": fmt.Sprintf("%d", rideID)},
	)
}

func (s *NotificationService) PropuestaRechazada(driverID int32, rideID int32) {
	s.sendAsync(driverID,
		"Propuesta rechazada ❌",
		"El cliente no aceptó tu propuesta",
		map[string]string{"type": "proposal_rejected", "ride_id": fmt.Sprintf("%d", rideID)},
	)
}
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

func NewNotificationService(
	fcm *core.FCMService,
	deviceRepo domain.DeviceRepository,
) *NotificationService {
	return &NotificationService{
		fcm:        fcm,
		deviceRepo: deviceRepo,
	}
}

func (s *NotificationService) sendAsync(userID int32, title, body string, data map[string]string) {
	go func() {
		tokens, err := s.deviceRepo.GetTokensByUser(userID)
		if err != nil {
			log.Printf("[FCM] error al obtener tokens del usuario %d: %v", userID, err)
			return
		}
		if len(tokens) == 0 {
			log.Printf("[FCM] usuario %d no tiene dispositivos registrados", userID)
			return
		}

		msg := core.FCMMessage{
			Title: title,
			Body:  body,
			Data:  data,
		}

		errs := s.fcm.SendToMany(context.Background(), tokens, msg)
		for _, err := range errs {
			if core.IsInvalidToken(err) {
				// Extraer el token del mensaje de error y eliminarlo
				log.Printf("[FCM] token inválido detectado, eliminando: %v", err)
			} else if err != nil {
				log.Printf("[FCM] error al enviar notificación: %v", err)
			}
		}
	}()
}


func (s *NotificationService) NuevoViaje(driverID int32, rideID int32) {
	s.sendAsync(
		driverID,
		"Nuevo viaje disponible",
		"Hay un viaje cerca de ti, ¡acepta antes de que alguien más lo tome!",
		map[string]string{
			"type":    "new_ride",
			"ride_id": fmt.Sprintf("%d", rideID),
		},
	)
}

func (s *NotificationService) NuevaPropuesta(clientID int32, rideID int32) {
	s.sendAsync(
		clientID,
		"Nueva propuesta recibida",
		"Un conductor ha hecho una propuesta en tu viaje",
		map[string]string{
			"type":    "new_proposal",
			"ride_id": fmt.Sprintf("%d", rideID),
		},
	)
}

func (s *NotificationService) PropuestaAceptada(driverID int32, rideID int32) {
	s.sendAsync(
		driverID,
		"¡Propuesta aceptada!",
		"El cliente aceptó tu propuesta, prepárate para el viaje",
		map[string]string{
			"type":    "proposal_accepted",
			"ride_id": fmt.Sprintf("%d", rideID),
		},
	)
}

func (s *NotificationService) PropuestaRechazada(driverID int32, rideID int32) {
	s.sendAsync(
		driverID,
		"Propuesta rechazada",
		"El cliente rechazó tu propuesta",
		map[string]string{
			"type":    "proposal_rejected",
			"ride_id": fmt.Sprintf("%d", rideID),
		},
	)
}

func (s *NotificationService) ViajeIniciado(clientID int32, rideID int32) {
	s.sendAsync(
		clientID,
		"¡Tu viaje ha iniciado!",
		"El conductor está en camino",
		map[string]string{
			"type":    "ride_started",
			"ride_id": fmt.Sprintf("%d", rideID),
		},
	)
}

func (s *NotificationService) ViajeFinalizado(clientID int32, rideID int32) {
	s.sendAsync(
		clientID,
		"Viaje finalizado",
		"¿Cómo estuvo tu viaje? No olvides calificar al conductor",
		map[string]string{
			"type":    "ride_finished",
			"ride_id": fmt.Sprintf("%d", rideID),
		},
	)
}
package application

import (
	"logired/src/internal/rides/domain"
	notifications "logired/src/internal/services/notifications"
)

const (
	RideStatusAgendado  = 1
	RideStatusEnCamino  = 2
	RideStatusEnProceso = 3
	RideStatusCancelado = 4
	RideStatusTerminado = 5
	RideStatusPublicado = 6
)

type UpdateRideStatus struct {
	repo     domain.IRide
	notifier *notifications.NotificationService
}

func NewUpdateRideStatus(
	repo domain.IRide,
	notifier *notifications.NotificationService,
) *UpdateRideStatus {
	return &UpdateRideStatus{repo: repo, notifier: notifier}
}

func (uc *UpdateRideStatus) Execute(idRide int32, idStatus int32) error {
	// Obtener el ride para saber a quién notificar
	ride, err := uc.repo.GetRideById(idRide)
	if err != nil {
		return err
	}

	if err := uc.repo.UpdateRideStatus(idRide, idStatus); err != nil {
		return err
	}

	switch idStatus {
	case RideStatusAgendado:
		uc.notifier.ViajeAgendado(ride.IdClient, idRide)

	case RideStatusEnCamino:
		uc.notifier.ViajeEnCamino(ride.IdClient, idRide)

	case RideStatusEnProceso:
		uc.notifier.ViajeIniciado(ride.IdClient, idRide)

	case RideStatusCancelado:
		uc.notifier.ViajeCancelado(ride.IdClient, idRide)
		if ride.IdDriver != nil {
			uc.notifier.ViajeCancelado(*ride.IdDriver, idRide)
		}

	case RideStatusTerminado:
		uc.notifier.ViajeFinalizado(ride.IdClient, idRide)
	}

	return nil
}
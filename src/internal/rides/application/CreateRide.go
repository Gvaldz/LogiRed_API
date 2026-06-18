package application

import (
	"log"
	notifications "logired/src/core/services/notifications"
	driverDomain "logired/src/internal/drivers/domain"
	"logired/src/internal/rides/domain"
	"logired/src/internal/rides/domain/entities"
)

type CreateRide struct {
	repo       domain.IRide
	driverRepo driverDomain.IDriver
	notifier   *notifications.NotificationService
}

func NewCreateRide(
	repo domain.IRide,
	driverRepo driverDomain.IDriver,
	notifier *notifications.NotificationService,
) *CreateRide {
	return &CreateRide{repo: repo, driverRepo: driverRepo, notifier: notifier}
}

func (uc *CreateRide) Execute(ride entities.Ride) error {
    if err := uc.repo.CreateRide(ride); err != nil {
        return err
    }

    go func() {
        drivers, err := uc.driverRepo.GetApprovedDrivers()
        if err != nil {
            log.Printf("[FCM] error al obtener conductores aprobados: %v", err)
            return
        }

        if len(drivers) == 0 {
            log.Printf("[FCM] no hay conductores aprobados")
            return
        }

        for _, driver := range drivers {
            uc.notifier.NuevoViaje(driver.IdUser, ride.IdRide)
        }

        log.Printf("[FCM] notificados %d conductores aprobados", len(drivers))
    }()

    return nil
}
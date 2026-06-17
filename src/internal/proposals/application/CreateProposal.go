package application

import (
	notifications "logired/src/core/services/notifications"
	"logired/src/internal/proposals/domain"
	"logired/src/internal/proposals/domain/entities"
	rideDomain "logired/src/internal/rides/domain"
)

type CreateProposal struct {
	repo     domain.IProposal
	rideRepo rideDomain.IRide // para obtener el IdClient del viaje
	notifier *notifications.NotificationService
}

func NewCreateProposal(
	repo domain.IProposal,
	rideRepo rideDomain.IRide,
	notifier *notifications.NotificationService,
) *CreateProposal {
	return &CreateProposal{repo: repo, rideRepo: rideRepo, notifier: notifier}
}

func (uc *CreateProposal) Execute(proposal entities.Proposal) error {
	if err := uc.repo.CreateProposal(proposal); err != nil {
		return err
	}

	go func() {
		ride, err := uc.rideRepo.GetRideById(proposal.IdRide)
		if err != nil {
			return
		}
		uc.notifier.NuevaPropuesta(ride.IdClient, proposal.IdRide)
	}()

	return nil
}

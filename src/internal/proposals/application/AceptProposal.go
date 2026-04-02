package application

import (
	"logired/src/internal/proposals/domain"
	notifications "logired/src/internal/services/notifications"
	rideDomain "logired/src/internal/rides/domain"
	"fmt"
)

const (
	ProposalStatusAceptada  = 1
	ProposalStatusEnProceso = 2
	ProposalStatusRechazada = 3
)

type AcceptProposal struct {
    repo     domain.IProposal
    rideRepo rideDomain.IRide  
    notifier *notifications.NotificationService
}

func NewAcceptProposal(
    repo domain.IProposal,
    rideRepo rideDomain.IRide, 
    notifier *notifications.NotificationService,
) *AcceptProposal {
    return &AcceptProposal{repo: repo, rideRepo: rideRepo, notifier: notifier}
}

func (uc *AcceptProposal) Execute(idProposal int32, idStatus int32) error {
    proposal, err := uc.repo.GetProposalById(idProposal)
    if err != nil {
        return err
    }

    if err := uc.repo.AcceptProposal(idProposal, idStatus); err != nil {
        return err
    }

    if idStatus == ProposalStatusAceptada {
        if err := uc.rideRepo.AssignDriver(proposal.IdRide, proposal.IdDriver); err != nil {
            return fmt.Errorf("error al asignar conductor al viaje: %w", err)
        }
        uc.notifier.PropuestaAceptada(proposal.IdDriver, proposal.IdRide)
    }

    if idStatus == ProposalStatusRechazada {
        uc.notifier.PropuestaRechazada(proposal.IdDriver, proposal.IdRide)
    }

    return nil
}
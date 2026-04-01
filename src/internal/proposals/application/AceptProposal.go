package application

import (
	"logired/src/internal/proposals/domain"
	notifications "logired/src/internal/services/notifications"
)

const (
	ProposalStatusAceptada  = 1
	ProposalStatusEnProceso = 2
	ProposalStatusRechazada = 3
)

type AcceptProposal struct {
	repo     domain.IProposal
	notifier *notifications.NotificationService
}

func NewAcceptProposal(
	repo domain.IProposal,
	notifier *notifications.NotificationService,
) *AcceptProposal {
	return &AcceptProposal{repo: repo, notifier: notifier}
}

func (uc *AcceptProposal) Execute(idProposal int32, idStatus int32) error {
	proposal, err := uc.repo.GetProposalById(idProposal)
	if err != nil {
		return err
	}

	if err := uc.repo.AcceptProposal(idProposal); err != nil {
		return err
	}

	switch idStatus {
	case ProposalStatusAceptada:
		uc.notifier.PropuestaAceptada(proposal.IdDriver, proposal.IdRide)
	case ProposalStatusRechazada:
		uc.notifier.PropuestaRechazada(proposal.IdDriver, proposal.IdRide)
	}

	return nil
}
package application

import "logired/src/internal/proposals/domain"

type AcceptProposal struct {
	repo domain.IProposal
}

func NewAcceptProposal(repo domain.IProposal) *AcceptProposal {
	return &AcceptProposal{repo: repo}
}

func (ap *AcceptProposal) Execute(idProposal int32) error {
	return ap.repo.AcceptProposal(idProposal)
}
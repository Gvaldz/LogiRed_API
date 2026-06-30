package application

import (
	"logired/src/internal/proposals/domain"
	"logired/src/internal/proposals/domain/entities"
)

type GetProposalsByDriver struct {
	repo domain.IProposal
}

func NewGetProposalsByDriver(repo domain.IProposal) *GetProposalsByDriver {
	return &GetProposalsByDriver{repo: repo}
}

func (uc *GetProposalsByDriver) Execute(idDriver int32) ([]entities.Proposal, error) {
	return uc.repo.GetProposalsByDriverId(idDriver)
}

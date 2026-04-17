package application

import (
    "logired/src/internal/proposals/domain"
    "logired/src/internal/proposals/domain/entities"
)

type GetProposalsByRide struct {
    repo domain.IProposal
}

func NewGetProposalsByRide(repo domain.IProposal) *GetProposalsByRide {
    return &GetProposalsByRide{repo: repo}
}

func (uc *GetProposalsByRide) Execute(idRide int32) ([]entities.Proposal, error) {
    return uc.repo.GetProposalsByRideId(idRide)
}
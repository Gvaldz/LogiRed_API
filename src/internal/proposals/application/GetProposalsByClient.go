package application

import (
	"logired/src/internal/proposals/domain"
	"logired/src/internal/proposals/domain/entities"
)

type GetProposalsByClient struct {
	repo domain.IProposal
}

func NewGetProposalsByClient(repo domain.IProposal) *GetProposalsByClient {
	return &GetProposalsByClient{repo: repo}
}

func (uc *GetProposalsByClient) Execute(idClient int32) ([]entities.Proposal, error) {
	return uc.repo.GetProposalsByClientId(idClient)
}

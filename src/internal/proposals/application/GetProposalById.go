package application

import (
	"logired/src/internal/proposals/domain"
	"logired/src/internal/proposals/domain/entities"
)

type GetProposalById struct {
	repo domain.IProposal
}

func NewGetProposalById(repo domain.IProposal) *GetProposalById {
	return &GetProposalById{repo: repo}
}

func (g *GetProposalById) Execute(idProposal int32) (entities.Proposal, error) {
	return g.repo.GetProposalById(idProposal)
}

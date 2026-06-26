package application

import (
	"logired/src/internal/proposals/domain"
)

type GetProposalById struct {
	repo domain.IProposal
}

func NewGetProposalById(repo domain.IProposal) *GetProposalById {
	return &GetProposalById{repo: repo}
}

func (g *GetProposalById) Execute(idProposal int32) (domain.ProposalDetail, error) {
	return g.repo.GetProposalDetailById(idProposal)
}

package application

import (
	"logired/src/internal/proposals/domain"
	"logired/src/internal/proposals/domain/entities"
)

type CreateProposal struct {
	repo domain.IProposal
}

func NewCreateProposal(repo domain.IProposal) *CreateProposal {
	return &CreateProposal{repo: repo}
}

func (cp *CreateProposal) Execute(proposal entities.Proposal) error {
	return cp.repo.CreateProposal(proposal)
}
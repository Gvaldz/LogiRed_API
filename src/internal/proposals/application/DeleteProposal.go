package application

import "logired/src/internal/proposals/domain"

type DeleteProposal struct {
	repo domain.IProposal
}

func NewDeleteProposal(repo domain.IProposal) *DeleteProposal {
	return &DeleteProposal{repo: repo}
}

func (dp *DeleteProposal) Execute(idProposal int32, idDriver int32) error {
	return dp.repo.DeleteProposal(idProposal, idDriver)
}
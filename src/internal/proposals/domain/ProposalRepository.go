package domain

import "logired/src/internal/proposals/domain/entities"

type IProposal interface {
	CreateProposal(proposal entities.Proposal) error
	GetProposalById(idProposal int32) (entities.Proposal, error) 
	AcceptProposal(idProposal int32, idStatus int32) error 
	DeleteProposal(idProposal int32, idDriver int32) error
}
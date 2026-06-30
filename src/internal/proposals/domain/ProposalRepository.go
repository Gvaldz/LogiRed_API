package domain

import "logired/src/internal/proposals/domain/entities"

type ProposalDetail struct {
	IdProposal      int32   `json:"idproposal"`
	Price           float32 `json:"price"`
	Comment         string  `json:"comment"`
	IdDriver        int32   `json:"iddriver"`
	IdRide          int32   `json:"idride"`
	IdStatus        int32   `json:"idproposalstatus"`
	Rating          float32 `json:"rating"`
	TripCount       int32   `json:"trip_count"`
	IdCar           *int32  `json:"idcar"`
	CarRegistration *string `json:"car_registration"`
	Brand           *string `json:"brand"`
	Model           *string `json:"model"`
	Color           *string `json:"color"`
	MaxCapacity     *int32  `json:"max_capacity"`
	FrontViewImage  *string `json:"front_view_image"`
	BackViewImage   *string `json:"back_view_image"`
	LeftViewImage   *string `json:"left_view_image"`
	RightViewImage  *string `json:"right_view_image"`
	PlatesImage     *string `json:"plates_image"`
	SpacesImage     *string `json:"space_image"`
}

type IProposal interface {
	CreateProposal(proposal entities.Proposal) error
	GetProposalById(idProposal int32) (entities.Proposal, error)
	GetProposalDetailById(idProposal int32) (ProposalDetail, error)
	GetProposalsByRideId(idRide int32) ([]entities.Proposal, error)
	AcceptProposal(idProposal int32, idStatus int32) error
	DeleteProposal(idProposal int32, idDriver int32) error
	GetProposalsByDriverId(idDriver int32) ([]entities.Proposal, error)
	GetProposalsByClientId(idClient int32) ([]entities.Proposal, error)
}
package application

import (
    "fmt"
    "logired/src/internal/proposals/domain"
    rideDomain "logired/src/internal/rides/domain"
)

const userTypeConductor = 2

type GetProposalById struct {
    repo     domain.IProposal
    rideRepo rideDomain.IRide
}

func NewGetProposalById(repo domain.IProposal, rideRepo rideDomain.IRide) *GetProposalById {
    return &GetProposalById{repo: repo, rideRepo: rideRepo}
}

func (g *GetProposalById) Execute(idProposal int32, requesterID int32, userType int) (domain.ProposalDetail, error) {
    detail, err := g.repo.GetProposalDetailById(idProposal)
    if err != nil {
        return domain.ProposalDetail{}, err
    }

    if userType == userTypeConductor {
        if detail.IdDriver != requesterID {
            return domain.ProposalDetail{}, fmt.Errorf("no tienes permiso para ver esta propuesta")
        }
    } else {
        ride, err := g.rideRepo.GetRideById(detail.IdRide)
        if err != nil {
            return domain.ProposalDetail{}, fmt.Errorf("viaje no encontrado")
        }
        if ride.IdClient != requesterID {
            return domain.ProposalDetail{}, fmt.Errorf("no tienes permiso para ver esta propuesta")
        }
    }

    return detail, nil
}

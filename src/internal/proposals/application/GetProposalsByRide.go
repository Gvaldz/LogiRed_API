package application

import (
    "fmt"
    "logired/src/internal/proposals/domain"
    "logired/src/internal/proposals/domain/entities"
    rideDomain "logired/src/internal/rides/domain"
)

type GetProposalsByRide struct {
    repo     domain.IProposal
    rideRepo rideDomain.IRide
}

func NewGetProposalsByRide(repo domain.IProposal, rideRepo rideDomain.IRide) *GetProposalsByRide {
    return &GetProposalsByRide{repo: repo, rideRepo: rideRepo}
}

func (uc *GetProposalsByRide) Execute(idRide int32, requesterID int32) ([]entities.Proposal, error) {
    ride, err := uc.rideRepo.GetRideById(idRide)
    if err != nil {
        return nil, fmt.Errorf("viaje no encontrado")
    }
    if ride.IdClient != requesterID {
        return nil, fmt.Errorf("solo el cliente dueño del viaje puede ver sus propuestas")
    }
    return uc.repo.GetProposalsByRideId(idRide)
}

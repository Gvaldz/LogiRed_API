package application

import (
	"logired/src/internal/rides/domain"
	"logired/src/internal/rides/domain/entities"
)

type GetRidesByClient struct {
	repo domain.IRide
}

func NewGetRidesByClient(repo domain.IRide) *GetRidesByClient {
	return &GetRidesByClient{repo: repo}
}

func (g *GetRidesByClient) Execute(idClient int32) ([]entities.Ride, error) {
	return g.repo.GetRidesByClientId(idClient)
}
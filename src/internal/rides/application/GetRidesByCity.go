package application

import (
	"logired/src/internal/rides/domain"
	"logired/src/internal/rides/domain/entities"
)

type GetRidesByCity struct {
	repo domain.IRide
}

func NewGetRidesByCity(repo domain.IRide) *GetRidesByCity {
	return &GetRidesByCity{repo: repo}
}

func (g *GetRidesByCity) Execute(city string) ([]entities.Ride, error) {
	return g.repo.GetRidesByCity(city)
}
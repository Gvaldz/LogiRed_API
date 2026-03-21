package application

import (
	"logired/src/internal/rides/domain"
	"logired/src/internal/rides/domain/entities"
)

type GetRideById struct {
	repo domain.IRide
}

func NewGetRideById(repo domain.IRide) *GetRideById {
	return &GetRideById{repo: repo}
}

func (g *GetRideById) Execute(idRide int32) (entities.Ride, error) {
	return g.repo.GetRideById(idRide)
}
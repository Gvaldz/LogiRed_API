package application

import (
	"logired/src/internal/rides/domain"
	"logired/src/internal/rides/domain/entities"
)

type GetRidesByDriver struct {
	repo domain.IRide
}

func NewGetRidesByDriver(repo domain.IRide) *GetRidesByDriver {
	return &GetRidesByDriver{repo: repo}
}

func (g *GetRidesByDriver) Execute(idDriver int32) ([]entities.Ride, error) {
	return g.repo.GetRidesByDriverId(idDriver)
}
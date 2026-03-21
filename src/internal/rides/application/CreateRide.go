package application

import (
	"logired/src/internal/rides/domain"
	"logired/src/internal/rides/domain/entities"
)

type CreateRide struct {
	repo domain.IRide
}

func NewCreateRide(repo domain.IRide) *CreateRide {
	return &CreateRide{repo: repo}
}

func (cr *CreateRide) Execute(ride entities.Ride) error {
	return cr.repo.CreateRide(ride)
}
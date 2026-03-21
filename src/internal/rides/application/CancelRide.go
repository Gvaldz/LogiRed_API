package application

import "logired/src/internal/rides/domain"

type CancelRide struct {
	repo domain.IRide
}

func NewCancelRide(repo domain.IRide) *CancelRide {
	return &CancelRide{repo: repo}
}

func (cr *CancelRide) Execute(idRide int32, idClient int32) error {
	return cr.repo.CancelRide(idRide, idClient)
}
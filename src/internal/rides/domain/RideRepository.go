package domain

import "logired/src/internal/rides/domain/entities"

type IRide interface {
	CreateRide(ride entities.Ride) error
	GetRidesByClientId(idClient int32) ([]entities.Ride, error)
	GetRideById(idRide int32) (entities.Ride, error)
	GetRidesByDriverId(idDriver int32) ([]entities.Ride, error)
	GetRidesByCity(city string) ([]entities.Ride, error)
	UpdateRideStatus(idRide int32, idStatus int32) error
}
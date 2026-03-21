package domain

import "logired/src/internal/cars/domain/entities"

type ICar interface {
	CreateCar(car entities.Car) error
	UpdateCar(car entities.Car) error
	GetCarsByDriverId(idDriver int32) ([]entities.Car, error)
	GetCarById(idCar int32, idDriver int32) (entities.Car, error)
	DeleteCar(idCar int32, idDriver int32) error
}
package application

import (
	"logired/src/internal/cars/domain"
	"logired/src/internal/cars/domain/entities"
)

type CreateCar struct {
	repo domain.ICar
}

func NewCreateCar(repo domain.ICar) *CreateCar {
	return &CreateCar{repo: repo}
}

func (uc *CreateCar) Execute(car entities.Car) error {
	return uc.repo.CreateCar(car)
}
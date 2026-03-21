package application

import (
	"logired/src/internal/cars/domain"
	"logired/src/internal/cars/domain/entities"
)

type UpdateCar struct {
	repo domain.ICar
}

func NewUpdateCar(repo domain.ICar) *UpdateCar {
	return &UpdateCar{repo: repo}
}

func (uc *UpdateCar) Execute(car entities.Car) error {
	return uc.repo.UpdateCar(car)
}
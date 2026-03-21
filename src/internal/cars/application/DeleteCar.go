package application

import (
	"logired/src/internal/cars/domain"
)

type DeleteCar struct {
	repo domain.ICar
}

func NewDeleteCar(repo domain.ICar) *DeleteCar {
	return &DeleteCar{repo: repo}
}

func (uc *DeleteCar) Execute(idCar int32, idDriver int32) error {
	return uc.repo.DeleteCar(idCar, idDriver)
}
package application

import (
	"logired/src/internal/cars/domain"
	"logired/src/internal/cars/domain/entities"
)

type GetCarById struct {
	repo domain.ICar
}

func NewGetCarById(repo domain.ICar) *GetCarById {
	return &GetCarById{repo: repo}
}

func (uc *GetCarById) Execute(idCar int32, idDriver int32) (entities.Car, error) {
	return uc.repo.GetCarById(idCar, idDriver)
}
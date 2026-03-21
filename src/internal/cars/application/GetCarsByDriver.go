package application

import (
	"logired/src/internal/cars/domain"
	"logired/src/internal/cars/domain/entities"
)

type GetCarsByDriver struct {
	repo domain.ICar
}

func NewGetCarsByDriver(repo domain.ICar) *GetCarsByDriver {
	return &GetCarsByDriver{repo: repo}
}

func (uc *GetCarsByDriver) Execute(idDriver int32) ([]entities.Car, error) {
	return uc.repo.GetCarsByDriverId(idDriver)
}
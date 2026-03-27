package application

import (
	"logired/src/internal/drivers/domain"
)

type GetDriversByCity struct {
	repo domain.IDriver
}

func NewGetDriversByCity(repo domain.IDriver) *GetDriversByCity {
	return &GetDriversByCity{repo: repo}
}

func (uc *GetDriversByCity) Execute(city string) ([]domain.DriverDetail, error) {
	return uc.repo.GetDriversByCity(city)
}
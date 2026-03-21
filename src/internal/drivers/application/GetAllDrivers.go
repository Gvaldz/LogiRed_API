package application

import (
	"logired/src/internal/drivers/domain"
)

type GetAllDrivers struct {
	repo domain.IDriver
}

func NewGetAllDrivers(repo domain.IDriver) *GetAllDrivers {
	return &GetAllDrivers{repo: repo}
}

func (uc *GetAllDrivers) Execute() ([]domain.DriverDetail, error) {
	return uc.repo.GetAll()
}
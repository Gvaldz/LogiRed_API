package application

import (
	"logired/src/internal/drivers/domain"
)

type GetDriverByID struct {
	repo domain.IDriver
}

func NewGetDriverByID(repo domain.IDriver) *GetDriverByID {
	return &GetDriverByID{repo: repo}
}

func (uc *GetDriverByID) Execute(driverID int32) (*domain.DriverDetail, error) {
	return uc.repo.GetByID(driverID)
}
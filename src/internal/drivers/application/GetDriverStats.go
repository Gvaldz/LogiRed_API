package application

import (
	"logired/src/internal/drivers/domain"
)

type GetDriverStats struct {
	repo domain.IDriver
}

func NewGetDriverStats(repo domain.IDriver) *GetDriverStats {
	return &GetDriverStats{repo: repo}
}

func (uc *GetDriverStats) Execute(driverID int32) (*domain.DriverStats, error) {
	return uc.repo.GetDriverStats(driverID)
}

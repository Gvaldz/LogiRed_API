package application

import (
	"logired/src/internal/drivers/domain"
)

type GetDriverByUser struct {
	repo domain.IDriver
}

func NewGetDriverByUser(repo domain.IDriver) *GetDriverByUser {
	return &GetDriverByUser{repo: repo}
}

func (uc *GetDriverByUser) Execute(userID int32) (*domain.DriverDetail, error) {
	return uc.repo.GetByUserID(userID)
}
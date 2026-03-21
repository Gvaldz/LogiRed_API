package application

import (
	"logired/src/internal/drivers/domain"
)

type DeleteDriver struct {
	repo domain.IDriver
}

func NewDeleteDriver(repo domain.IDriver) *DeleteDriver {
	return &DeleteDriver{repo: repo}
}

func (uc *DeleteDriver) Execute(userID int32) error {
	return uc.repo.Delete(userID)
}
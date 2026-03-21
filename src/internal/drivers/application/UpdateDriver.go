package application

import (
	"logired/src/internal/drivers/domain"
	"logired/src/internal/drivers/domain/entities"
)

type UpdateDriver struct {
	repo domain.IDriver
}

func NewUpdateDriver(repo domain.IDriver) *UpdateDriver {
	return &UpdateDriver{repo: repo}
}

func (uc *UpdateDriver) Execute(driver entities.Driver) error {
	return uc.repo.Update(driver)
}
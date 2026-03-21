package application

import (
	"logired/src/internal/drivers/domain"
	"logired/src/internal/drivers/domain/entities"
)

type CreateDriver struct {
	repo domain.IDriver
}

func NewCreateDriver(repo domain.IDriver) *CreateDriver {
	return &CreateDriver{repo: repo}
}

func (uc *CreateDriver) Execute(driver entities.Driver) error {
	return uc.repo.Create(driver)
}
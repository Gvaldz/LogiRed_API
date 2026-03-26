package application

import "logired/src/internal/drivers/domain"

type UpdateDriverProfile struct {
    repo domain.IDriver
}

func NewUpdateDriverProfile(repo domain.IDriver) *UpdateDriverProfile {
    return &UpdateDriverProfile{repo: repo}
}

func (uc *UpdateDriverProfile) Execute(idUser int32, citywork string) error {
    return uc.repo.UpdateCitywork(idUser, citywork)
}
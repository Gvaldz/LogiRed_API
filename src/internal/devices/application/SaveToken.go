package application

import "logired/src/internal/devices/domain"

type SaveDeviceToken struct {
	repo domain.DeviceRepository
}

func NewSaveDeviceToken(repo domain.DeviceRepository) *SaveDeviceToken {
	return &SaveDeviceToken{repo: repo}
}

func (uc *SaveDeviceToken) Execute(userID int32, token, deviceName string) error {
	return uc.repo.SaveToken(userID, token, deviceName)
}
package domain

type DeviceRepository interface {
	SaveToken(userID int32, token, deviceName string) error
	GetTokensByUser(userID int32) ([]string, error)
	DeleteToken(token string) error
}
package entities

type Device struct {
	ID         int32  `json:"id"`
	IdUser     int32  `json:"iduser"`
	FCMToken   string `json:"fcm_token"`
	DeviceName string `json:"device_name"`
}
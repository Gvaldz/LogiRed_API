package dto

type DeviceTokenRequest struct {
    Token      string `json:"fcm_token" binding:"required"`
    DeviceName string `json:"device_name"`
}
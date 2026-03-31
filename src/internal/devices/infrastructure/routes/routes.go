package routes

import (
	"logired/src/internal/devices/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

type DeviceRoutes struct {
	controller     *controllers.DeviceTokenController
	authMiddleware gin.HandlerFunc
}

func NewDeviceRoutes(
	controller *controllers.DeviceTokenController,
	authMiddleware gin.HandlerFunc,
) *DeviceRoutes {
	return &DeviceRoutes{
		controller:     controller,
		authMiddleware: authMiddleware,
	}
}

func (r *DeviceRoutes) AttachRoutes(router *gin.Engine) {
	devices := router.Group("/devices")
	devices.Use(r.authMiddleware)
	{
		devices.PUT("/token", r.controller.Register)
	}
}
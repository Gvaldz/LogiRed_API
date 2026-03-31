package dependences

import (
	"database/sql"
	"logired/src/internal/devices/application"
	"logired/src/internal/devices/infrastructure/controllers"
	"logired/src/internal/devices/infrastructure/repositories"
	"logired/src/internal/devices/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

type DeviceDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
}

func NewDeviceDependencies(db *sql.DB, authMiddleware gin.HandlerFunc) *DeviceDependencies {
	return &DeviceDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
	}
}

func (d *DeviceDependencies) GetRoutes() *routes.DeviceRoutes {
	repo       := repositories.NewDeviceRepo(d.DB)
	useCase    := application.NewSaveDeviceToken(repo)
	controller := controllers.NewDeviceTokenController(useCase)

	return routes.NewDeviceRoutes(controller, d.AuthMiddleware)
}
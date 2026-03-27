package dependencies

import (
	"database/sql"
	"logired/src/internal/drivers/application"
	"logired/src/internal/drivers/infrastructure/controllers"
	"logired/src/internal/drivers/infrastructure/repositories"
	"logired/src/internal/drivers/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

type DriverDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
}

func NewDriverDependencies(db *sql.DB, authMiddleware gin.HandlerFunc) *DriverDependencies {
	return &DriverDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
	}
}

func (d *DriverDependencies) GetRoutes() *routes.DriverRoutes {
	driverRepo := repositories.NewDriverRepo(d.DB)

	getByCityUseCase := application.NewGetDriversByCity(driverRepo)

	getByCityController := controllers.NewGetRideByCityController(getByCityUseCase)


	return routes.NewDriverRoutes(
		getByCityController,
		d.AuthMiddleware,
	)
}
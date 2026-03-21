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

	createUseCase := application.NewCreateDriver(driverRepo)
	getByUserUseCase := application.NewGetDriverByUser(driverRepo)
	getByIDUseCase := application.NewGetDriverByID(driverRepo)
	getAllUseCase := application.NewGetAllDrivers(driverRepo)
	updateUseCase := application.NewUpdateDriver(driverRepo)
	deleteUseCase := application.NewDeleteDriver(driverRepo)

	createController := controllers.NewCreateDriverController(createUseCase)
	getController := controllers.NewGetDriverController(getByUserUseCase, getByIDUseCase)
	getAllController := controllers.NewGetAllDriversController(getAllUseCase)
	updateController := controllers.NewUpdateDriverController(updateUseCase)
	deleteController := controllers.NewDeleteDriverController(deleteUseCase)

	return routes.NewDriverRoutes(
		createController,
		getController,
		getAllController,
		updateController,
		deleteController,
		d.AuthMiddleware,
	)
}
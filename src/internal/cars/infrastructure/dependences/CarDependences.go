package dependencies

import (
	"database/sql"
	"logired/src/internal/cars/application"
	"logired/src/internal/cars/infrastructure/controllers"
	"logired/src/internal/cars/infrastructure/repositories"
	"logired/src/internal/cars/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

type CarDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
}

func NewCarDependencies(db *sql.DB, authMiddleware gin.HandlerFunc) *CarDependencies {
	return &CarDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
	}
}

func (d *CarDependencies) GetRoutes() *routes.CarRoutes {
	carRepo := repositories.NewCarRepo(d.DB)

	createCarUseCase := application.NewCreateCar(carRepo)
	updateCarUseCase := application.NewUpdateCar(carRepo)
	getCarsByDriverUseCase := application.NewGetCarsByDriver(carRepo)
	getCarByIdUseCase := application.NewGetCarById(carRepo)
	deleteCarUseCase := application.NewDeleteCar(carRepo)

	createCarController := controllers.NewCreateCarController(createCarUseCase)
	updateCarController := controllers.NewUpdateCarController(updateCarUseCase)
	getCarsByDriverController := controllers.NewGetCarsByDriverController(getCarsByDriverUseCase)
	getCarByIdController := controllers.NewGetCarByIdController(getCarByIdUseCase)
	deleteCarController := controllers.NewDeleteCarController(deleteCarUseCase)

	return routes.NewCarRoutes(
		createCarController,
		updateCarController,
		getCarsByDriverController,
		getCarByIdController,
		deleteCarController,
		d.AuthMiddleware,
	)
}
package dependencies

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"logired/src/internal/cars/application"
	"logired/src/internal/cars/infrastructure/routes"
	"logired/src/internal/cars/infrastructure/controllers"
	"logired/src/internal/cars/infrastructure/repositories"
	storageService 		"logired/src/core/services/storage"
)

type CarDependencies struct {
	DB             		*sql.DB
	AuthMiddleware 		gin.HandlerFunc
	StorageService	 	storageService.StorageService

}

func NewCarDependencies(
	db				 *sql.DB,
	authMiddleware 	 gin.HandlerFunc, 	
	storageService 	 storageService.StorageService,
) *CarDependencies {
	return &CarDependencies{
		DB:              db,
		AuthMiddleware:  authMiddleware,
		StorageService:  storageService,
	}
}

func (d *CarDependencies) GetRoutes() *routes.CarRoutes {
	carRepo := repositories.NewCarRepo(d.DB)

	createCarUseCase := application.NewCreateCar(carRepo)
	updateCarUseCase := application.NewUpdateCar(carRepo)
	getCarsByDriverUseCase := application.NewGetCarsByDriver(carRepo)
	getCarByIdUseCase := application.NewGetCarById(carRepo)
	deleteCarUseCase := application.NewDeleteCar(carRepo)


	createCarController := controllers.NewCreateCarController(createCarUseCase, d.StorageService)
	updateCarController := controllers.NewUpdateCarController(updateCarUseCase, d.StorageService)
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
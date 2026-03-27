package dependencies

import (
	"database/sql"
	"logired/src/internal/rides/application"
	"logired/src/internal/rides/infrastructure/controllers"
	"logired/src/internal/rides/infrastructure/repositories"
	"logired/src/internal/rides/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

type RideDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
}

func NewRideDependencies(db *sql.DB, authMiddleware gin.HandlerFunc) *RideDependencies {
	return &RideDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
	}
}

func (d *RideDependencies) GetRoutes() *routes.RideRoutes {
	rideRepo := repositories.NewRideRepo(d.DB)

	createRideUseCase := application.NewCreateRide(rideRepo)
	cancelRideUseCase := application.NewCancelRide(rideRepo)
	getRidesByClientUseCase := application.NewGetRidesByClient(rideRepo)
	getRideByIdUseCase := application.NewGetRideById(rideRepo)
	getRidesByDriverUseCase := application.NewGetRidesByDriver(rideRepo)
	getRidesByCityUseCase := application.NewGetRidesByCity(rideRepo)

	createRideController := controllers.NewCreateRideController(createRideUseCase)
	cancelRideController := controllers.NewCancelRideController(cancelRideUseCase)
	getRidesByClientController := controllers.NewGetRidesByClientController(getRidesByClientUseCase)
	getRideByIdController := controllers.NewGetRideByIdController(getRideByIdUseCase)
	getRidesByDriverController := controllers.NewGetRidesByDriverController(getRidesByDriverUseCase)
	getRidesByCityController := controllers.NewGetRideByCityController(getRidesByCityUseCase)

	return routes.NewRideRoutes(
		createRideController,
		cancelRideController,
		getRidesByClientController,
		getRideByIdController,
		getRidesByDriverController,
		getRidesByCityController,
		d.AuthMiddleware,
	)
}
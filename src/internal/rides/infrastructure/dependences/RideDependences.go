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
	getRidesByClientUseCase := application.NewGetRidesByClient(rideRepo)
	getRideByIdUseCase := application.NewGetRideById(rideRepo)
	getRidesByDriverUseCase := application.NewGetRidesByDriver(rideRepo)
	getRidesByCityUseCase := application.NewGetRidesByCity(rideRepo)
	updateRideStatusUseCase := application.NewUpdateRideStatus(rideRepo)

	createRideController := controllers.NewCreateRideController(createRideUseCase)
	getRidesByClientController := controllers.NewGetRidesByClientController(getRidesByClientUseCase)
	getRideByIdController := controllers.NewGetRideByIdController(getRideByIdUseCase)
	getRidesByDriverController := controllers.NewGetRidesByDriverController(getRidesByDriverUseCase)
	getRidesByCityController := controllers.NewGetRideByCityController(getRidesByCityUseCase)
	updateRideStatusController := controllers.NewUpdateRideStatusController(updateRideStatusUseCase)

	return routes.NewRideRoutes(
		createRideController,
		getRidesByClientController,
		getRideByIdController,
		getRidesByDriverController,
		getRidesByCityController,
		updateRideStatusController,
		d.AuthMiddleware,
	)
}
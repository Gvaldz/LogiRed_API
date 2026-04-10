package dependencies

import (
	"database/sql"
	"logired/src/internal/rides/application"
	"logired/src/internal/rides/infrastructure/controllers"
	"logired/src/internal/rides/infrastructure/repositories"
	"logired/src/internal/rides/infrastructure/routes"
	driverRepositories "logired/src/internal/drivers/infrastructure/repositories"
	notifications "logired/src/internal/services/notifications"

	"github.com/gin-gonic/gin"
)

type RideDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
	Notifier       *notifications.NotificationService
}

func NewRideDependencies(
	db *sql.DB,
	authMiddleware gin.HandlerFunc,
	notifier *notifications.NotificationService,
) *RideDependencies {
	return &RideDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
		Notifier:       notifier,
	}
}

func (d *RideDependencies) GetRoutes() *routes.RideRoutes {
	rideRepo   := repositories.NewRideRepo(d.DB)
	driverRepo := driverRepositories.NewDriverRepo(d.DB)

	createRideUseCase             := application.NewCreateRide(rideRepo, driverRepo, d.Notifier)
	getRidesByClientUseCase       := application.NewGetRidesByClient(rideRepo)
	getRideByIdUseCase            := application.NewGetRideById(rideRepo)
	getRidesByDriverUseCase       := application.NewGetRidesByDriver(rideRepo)
	getRidesByCityUseCase         := application.NewGetRidesByCity(rideRepo)
	getRidesHistoryUseCase        := application.NewGetRidesHistory(rideRepo)
	getRidesByStatusUseCase    	  := application.NewGetRidesByStatus(rideRepo)
	updateRideStatusUseCase    	  := application.NewUpdateRideStatus(rideRepo, d.Notifier)

	createRideController          := controllers.NewCreateRideController(createRideUseCase)
	getRidesByClientController    := controllers.NewGetRidesByClientController(getRidesByClientUseCase)
	getRideByIdController         := controllers.NewGetRideByIdController(getRideByIdUseCase)
	getRidesByDriverController    := controllers.NewGetRidesByDriverController(getRidesByDriverUseCase)
	getRidesByCityController      := controllers.NewGetRideByCityController(getRidesByCityUseCase)
	getRidesByStatusController	  := controllers.NewGetRidesByStatusController(getRidesByStatusUseCase)
	getRidesHistoryController     := controllers.NewGetRidesHistoryController(getRidesHistoryUseCase)
	updateRideStatusController    := controllers.NewUpdateRideStatusController(updateRideStatusUseCase)

	return routes.NewRideRoutes(
		createRideController,
		getRidesByClientController,
		getRideByIdController,
		getRidesByDriverController,
		getRidesByCityController,
		getRidesHistoryController,
		getRidesByStatusController,
		updateRideStatusController,
		d.AuthMiddleware,
	)
}
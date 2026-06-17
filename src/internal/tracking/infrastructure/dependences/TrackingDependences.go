package dependences

import (
	"database/sql"

	tokenServiceDomain "logired/src/core/services/auth/domain"
	rideRepo "logired/src/internal/rides/infrastructure/repositories"
	"logired/src/internal/tracking/infrastructure"
	"logired/src/internal/tracking/infrastructure/controllers"
	"logired/src/internal/tracking/infrastructure/routes"
)

type TrackingDependencies struct {
	routes *routes.TrackingRoutes
}

func NewTrackingDependencies(
	db *sql.DB,
	tokenService tokenServiceDomain.TokenService,
) *TrackingDependencies {
	rideRepository := rideRepo.NewRideRepo(db)

	hub := infrastructure.NewHub()

	trackingController := controllers.NewTrackingController(
		hub,
		tokenService,
		rideRepository,
	)

	trackingRoutes := routes.NewTrackingRoutes(trackingController)

	return &TrackingDependencies{routes: trackingRoutes}
}

func (d *TrackingDependencies) GetRoutes() *routes.TrackingRoutes {
	return d.routes
}

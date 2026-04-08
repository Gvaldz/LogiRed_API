package dependences

import (
    "database/sql"

    rideRepo "logired/src/internal/rides/infrastructure/repositories"
    tokenServiceDomain "logired/src/internal/services/auth/domain"
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
    // Reusamos el mismo repo de rides para validar permisos
    rideRepository := rideRepo.NewRideRepo(db)

    // El hub es único para toda la app (estado en memoria de suscripciones)
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
package cmd

import (
	"log"
	"logired/src/core"
	carsDeps "logired/src/internal/cars/infrastructure/dependences"
	devicesDeps "logired/src/internal/devices/infrastructure/dependences"
	deviceRepo "logired/src/internal/devices/infrastructure/repositories"
	driversDeps "logired/src/internal/drivers/infrastructure/dependences"
	driverRepo "logired/src/internal/drivers/infrastructure/repositories"
	paymentDeps "logired/src/internal/payments/infrastructure/dependences"
	proposalDeps "logired/src/internal/proposals/infrastructure/dependences"
	reviewDeps "logired/src/internal/reviews/infrastructure/dependences"
	ridesDeps "logired/src/internal/rides/infrastructure/dependences"
	loginDeps "logired/src/internal/services/auth/infrastructure"
	notifications "logired/src/internal/services/notifications"
	trackingDeps "logired/src/internal/tracking/infrastructure/dependences"
	usersDeps "logired/src/internal/users/infrastructure"
	"logired/src/server"
	"logired/src/server/middleware"
)

func Init() {
	cfg := core.LoadConfig()

	db, err := core.ConnectDB()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	hasher := core.NewBcryptHasher(cfg.BcryptCost)
	tokenService := core.NewJWTService()
	fcmService := core.NewFCMService(cfg.FCMProjectID, cfg.FCMCredentialsFile)

	userRepo := usersDeps.NewUserRepository(db)
	authRepo := loginDeps.NewAuthRepository(db)
	driverRepo := driverRepo.NewDriverRepo(db)
	deviceRepo := deviceRepo.NewDeviceRepo(db)

	notifier := notifications.NewNotificationService(fcmService, deviceRepo)
	authMiddleware := middleware.AuthMiddleware(tokenService, userRepo)

	devicesRoutes := devicesDeps.NewDeviceDependencies(db, authMiddleware).GetRoutes()
	carsRoutes := carsDeps.NewCarDependencies(db, authMiddleware).GetRoutes()
	ridesRoutes := ridesDeps.NewRideDependencies(db, authMiddleware, notifier).GetRoutes()
	proposalRoutes := proposalDeps.NewProposalDependencies(db, authMiddleware, notifier).GetRoutes()
	reviewRoutes := reviewDeps.NewReviewDependencies(db, authMiddleware).GetRoutes()
	driversRoutes := driversDeps.NewDriverDependencies(db, authMiddleware).GetRoutes()
	trackingRoutes := trackingDeps.NewTrackingDependencies(db, tokenService).GetRoutes()
	paymentRoutes := paymentDeps.NewPaymentDependencies(db, authMiddleware, cfg.StripeSecretKey).GetRoutes()

	userRoutes := usersDeps.NewUserDependencies(
		db, hasher, tokenService, authRepo, userRepo, driverRepo,
	).GetRoutes()

	authRoutes := loginDeps.NewAuthDependencies(db, hasher, userRepo).GetRoutes()

	server.Run(
		authRoutes,
		userRoutes,
		devicesRoutes,
		carsRoutes,
		ridesRoutes,
		proposalRoutes,
		reviewRoutes,
		driversRoutes,
		trackingRoutes,
		paymentRoutes,
	)
}

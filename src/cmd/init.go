package cmd

import (
	"log"
	"logired/src/core"
	"logired/src/server"
	"logired/src/server/middleware"
	driverRepo   "logired/src/internal/drivers/infrastructure/repositories"
//	deviceRepo   "logired/src/internal/devices/infrastructure/repositories"
	loginDeps    "logired/src/internal/services/auth/infrastructure"
	usersDeps    "logired/src/internal/users/infrastructure"
	carsDeps     "logired/src/internal/cars/infrastructure/dependences"
	ridesDeps    "logired/src/internal/rides/infrastructure/dependences"
	proposalDeps "logired/src/internal/proposals/infrastructure/dependences"
	reviewDeps   "logired/src/internal/reviews/infrastructure/dependences"
	driversDeps  "logired/src/internal/drivers/infrastructure/dependences"
//	devicesDeps  "logired/src/internal/devices/infrastructure/dependences"
//	notifications "logired/src/internal/services/notifications"
)

func Init() {
	// ── Cargar configuración ───────────────────────────────────────
	cfg := core.LoadConfig()

	db, err := core.ConnectDB()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	// ── Servicios core ─────────────────────────────────────────────
	hasher       := core.NewBcryptHasher(cfg.BcryptCost)
	tokenService := core.NewJWTService()
//	fcmService   := core.NewFCMService(cfg.FCMProjectID, cfg.FCMCredentialsFile)

	// ── Repositorios ───────────────────────────────────────────────
	userRepo   := usersDeps.NewUserRepository(db)
	authRepo   := loginDeps.NewAuthRepository(db)
	driverRepo := driverRepo.NewDriverRepo(db)
//	deviceRepo := deviceRepo.NewDeviceRepo(db)

	// ── Servicios ──────────────────────────────────────────────────
//	notifier      := notifications.NewNotificationService(fcmService, deviceRepo)
	authMiddleware := middleware.AuthMiddleware(tokenService, userRepo)

	// ── Dependencias y rutas ───────────────────────────────────────
//	devicesRoutes  := devicesDeps.NewDeviceDependencies(db, authMiddleware).GetRoutes()
	carsRoutes     := carsDeps.NewCarDependencies(db, authMiddleware).GetRoutes()
	ridesRoutes    := ridesDeps.NewRideDependencies(db, authMiddleware).GetRoutes()
	proposalRoutes := proposalDeps.NewProposalDependencies(db, authMiddleware).GetRoutes()
	reviewRoutes   := reviewDeps.NewReviewDependencies(db, authMiddleware).GetRoutes()
	driversRoutes  := driversDeps.NewDriverDependencies(db, authMiddleware).GetRoutes()

	userRoutes := usersDeps.NewUserDependencies(
		db, hasher, tokenService, authRepo, userRepo, driverRepo,
	).GetRoutes()

	authRoutes := loginDeps.NewAuthDependencies(db, hasher, userRepo).GetRoutes()

	server.Run(
		authRoutes,
		userRoutes,
		carsRoutes,
		ridesRoutes,
		proposalRoutes,
		reviewRoutes,
		driversRoutes,
	)
}
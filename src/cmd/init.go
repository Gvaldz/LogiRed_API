package cmd

import (
    "log"
    "logired/src/core"
    loginDeps        "logired/src/core/services/auth/infrastructure"
    notifications    "logired/src/core/services/notifications"
    carsDeps         "logired/src/internal/cars/infrastructure/dependences"
    devicesDeps      "logired/src/internal/devices/infrastructure/dependences"
    deviceRepo       "logired/src/internal/devices/infrastructure/repositories"
    driverRepo       "logired/src/internal/drivers/infrastructure/repositories"
    paymentDeps      "logired/src/internal/payments/infrastructure/dependences"
    proposalDeps     "logired/src/internal/proposals/infrastructure/dependences"
    reviewDeps       "logired/src/internal/reviews/infrastructure/dependences"
    ridesDeps        "logired/src/internal/rides/infrastructure/dependences"
    trackingDeps     "logired/src/internal/tracking/infrastructure/dependences"
    usersDeps        "logired/src/internal/users/infrastructure"
    storage          "logired/src/core/services/storage"
    "logired/src/server"
    "logired/src/server/middleware"
    
    carRepos         "logired/src/internal/cars/infrastructure/repositories"
    documentRepos    "logired/src/internal/documents/infrastructure/repositories"
)

func Init() {
    cfg := core.LoadConfig()

    storageService := storage.InitStorage()

    db, err := core.ConnectDB()
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

    hasher := core.NewBcryptHasher(cfg.BcryptCost)
    tokenService := core.NewJWTService()
    fcmService := core.NewFCMService(cfg.FCMProjectID, cfg.FCMCredentialsFile)

    // Instanciamos todos los repositorios principales
    userRepo := usersDeps.NewUserRepository(db)
    authRepo := loginDeps.NewAuthRepository(db)
    driverRepo := driverRepo.NewDriverRepo(db)
    deviceRepo := deviceRepo.NewDeviceRepo(db)
    
    // Instanciamos los repositorios de carros y documentos a nivel raíz
    carRepo := carRepos.NewCarRepo(db)
    docRepo := documentRepos.NewDocumentRepo(db)

    notifier := notifications.NewNotificationService(fcmService, deviceRepo)
    authMiddleware := middleware.AuthMiddleware(tokenService, userRepo)
    driverApprovalMiddleware := middleware.RequireApprovedDriver(authRepo)
    driverOnlyApprovedMiddleware := middleware.RequireOnlyApprovedDriver(authRepo)
    protectedMiddleware := middleware.Chain(authMiddleware, driverApprovalMiddleware)

    devicesRoutes := devicesDeps.NewDeviceDependencies(db, protectedMiddleware).GetRoutes()
    carsRoutes := carsDeps.NewCarDependencies(db, protectedMiddleware, storageService).GetRoutes()
    ridesRoutes := ridesDeps.NewRideDependencies(db, authMiddleware, driverOnlyApprovedMiddleware, notifier).GetRoutes()
    proposalRoutes := proposalDeps.NewProposalDependencies(db, authMiddleware, driverOnlyApprovedMiddleware, notifier).GetRoutes()
    reviewRoutes := reviewDeps.NewReviewDependencies(db, protectedMiddleware).GetRoutes()
    trackingRoutes := trackingDeps.NewTrackingDependencies(db, tokenService).GetRoutes()
    paymentRoutes := paymentDeps.NewPaymentDependencies(db, protectedMiddleware, cfg.StripeSecretKey).GetRoutes()

    userRoutes := usersDeps.NewUserDependencies(
        db, hasher, tokenService, authRepo, userRepo, driverRepo, carRepo, docRepo, storageService,
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
        trackingRoutes,
        paymentRoutes,
    )
}

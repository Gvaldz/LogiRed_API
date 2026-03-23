package cmd

import (
	"log"
	"logired/src/core"
	"logired/src/server"
	"logired/src/server/middleware"
	driverRepo 		"logired/src/internal/drivers/infrastructure/repositories"
	loginDeps 		"logired/src/internal/services/auth/infrastructure"
	usersDeps	 	"logired/src/internal/users/infrastructure"
	carsDeps 		"logired/src/internal/cars/infrastructure/dependences"
	ridesDeps 		"logired/src/internal/rides/infrastructure/dependences"
	proposalDeps 	"logired/src/internal/proposals/infrastructure/dependences"
	reviewDeps 		"logired/src/internal/reviews/infrastructure/dependences"
	driversDeps 	"logired/src/internal/drivers/infrastructure/dependences"
)

func Init() {
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	hasher := core.NewBcryptHasher(14)
	tokenService := core.NewJWTService()

	userRepo := usersDeps.NewUserRepository(db)
	authRepo := loginDeps.NewAuthRepository(db)
	
	authMiddleware := middleware.AuthMiddleware(tokenService, userRepo)
	carsDeps := carsDeps.NewCarDependencies(db, authMiddleware)
	carsRoutes := carsDeps.GetRoutes()

	ridesDeps := ridesDeps.NewRideDependencies(db, authMiddleware)
	ridesRoutes := ridesDeps.GetRoutes()

	proposalDeps := proposalDeps.NewProposalDependencies(db, authMiddleware)
	proposalRoutes := proposalDeps.GetRoutes()

	reviewDeps := reviewDeps.NewReviewDependencies(db, authMiddleware)
	reviewRoutes := reviewDeps.GetRoutes()

	driverRepo := driverRepo.NewDriverRepo(db)
	driversDeps := driversDeps.NewDriverDependencies(db, authMiddleware)
	driversRoutes := driversDeps.GetRoutes()


	userDependencies := usersDeps.NewUserDependencies(
		db,
		hasher,
		tokenService,
		authRepo,
		userRepo,
		driverRepo,
	)
	userRoutes := userDependencies.GetRoutes()

	authDependencies := loginDeps.NewAuthDependencies(db, hasher, userRepo)
	authRoutes := authDependencies.GetRoutes()

	server.Run(authRoutes, userRoutes, carsRoutes, ridesRoutes, proposalRoutes, reviewRoutes, driversRoutes)
}

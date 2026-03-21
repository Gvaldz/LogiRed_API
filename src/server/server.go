package server

import (
	loginRouters	 "logired/src/internal/services/auth/infrastructure"
	userRouters 	 "logired/src/internal/users/infrastructure"
	carsRoutes 		 "logired/src/internal/cars/infrastructure/routes"
	ridesRoutes  	 "logired/src/internal/rides/infrastructure/routes"
	proposalRoutes 	 "logired/src/internal/proposals/infrastructure/routes"
	reviewRoutes 	 "logired/src/internal/reviews/infrastructure/routes"
	driversRoutes 	 "logired/src/internal/drivers/infrastructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(
	authRoutes *loginRouters.AuthRoutes,
	userRoutes *userRouters.UserRoutes,
	carsRoutes *carsRoutes.CarRoutes,
	ridesRoutes *ridesRoutes.RideRoutes,
	proposalRoutes *proposalRoutes.ProposalRoutes,
	reviewRoutes *reviewRoutes.ReviewRoutes,
	driversRoutes *driversRoutes.DriverRoutes,
) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	authRoutes.AttachRoutes(r)
	userRoutes.AttachRoutes(r)
	carsRoutes.AttachRoutes(r)
	ridesRoutes.AttachRoutes(r)
	proposalRoutes.AttachRoutes(r)
	reviewRoutes.AttachRoutes(r)
	driversRoutes.AttachRoutes(r)

	r.Run(":8080")
}

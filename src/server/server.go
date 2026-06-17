package server

import (
	"time"

	loginRouters "logired/src/core/services/auth/infrastructure"
	carsRoutes "logired/src/internal/cars/infrastructure/routes"
	devicesRoutes "logired/src/internal/devices/infrastructure/routes"
	driversRoutes "logired/src/internal/drivers/infrastructure/routes"
	paymentRoutes "logired/src/internal/payments/infrastructure/routes"
	proposalRoutes "logired/src/internal/proposals/infrastructure/routes"
	reviewRoutes "logired/src/internal/reviews/infrastructure/routes"
	ridesRoutes "logired/src/internal/rides/infrastructure/routes"
	trackingRoutes "logired/src/internal/tracking/infrastructure/routes"
	userRouters "logired/src/internal/users/infrastructure"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"logired/src/server/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "logired/docs"
)

func Run(
	authRoutes *loginRouters.AuthRoutes,
	userRoutes *userRouters.UserRoutes,
	devicesRoutes *devicesRoutes.DeviceRoutes,
	carsRoutes *carsRoutes.CarRoutes,
	ridesRoutes *ridesRoutes.RideRoutes,
	proposalRoutes *proposalRoutes.ProposalRoutes,
	reviewRoutes *reviewRoutes.ReviewRoutes,
	driversRoutes *driversRoutes.DriverRoutes,
	trackingRoutes *trackingRoutes.TrackingRoutes,
	paymentRoutes *paymentRoutes.PaymentRoutes,
) {

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	globalLimiter := middleware.NewRateLimiter(20, 40)
	r.Use(globalLimiter.Middleware())

	r.Use(middleware.MaxBodySize(20 * 1024 * 1024))

	r.Use(middleware.Timeout(30 * time.Second))

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	authGroup := r.Group("/")
	authGroup.Use(middleware.StrictRateLimit(5.0/60.0, 5))
	authRoutes.AttachRoutes(r)

	userRoutes.AttachRoutes(r)
	devicesRoutes.AttachRoutes(r)
	carsRoutes.AttachRoutes(r)
	ridesRoutes.AttachRoutes(r)
	proposalRoutes.AttachRoutes(r)
	reviewRoutes.AttachRoutes(r)
	driversRoutes.AttachRoutes(r)
	trackingRoutes.AttachRoutes(r)
	paymentRoutes.AttachRoutes(r)

	r.Run(":8081")
}

package routes

import (
	"logired/src/internal/drivers/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type DriverRoutes struct {
	getStatsController *controllers.GetDriverStatsController
	authMiddleware     gin.HandlerFunc
}

func NewDriverRoutes(
	getStats *controllers.GetDriverStatsController,
	authMiddleware gin.HandlerFunc,
) *DriverRoutes {
	return &DriverRoutes{
		getStatsController: getStats,
		authMiddleware:     authMiddleware,
	}
}

func (r *DriverRoutes) AttachRoutes(router *gin.Engine) {
	driversGroup := router.Group("/drivers")
	driversGroup.Use(r.authMiddleware)

	driversGroup.GET("/:id/stats", r.getStatsController.GetStats)
}
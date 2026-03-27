package routes

import (
	"logired/src/internal/drivers/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type DriverRoutes struct {
	getByCityController *controllers.GetDriversByCityController
	authMiddleware   gin.HandlerFunc
}

func NewDriverRoutes(
	getByCity *controllers.GetDriversByCityController,
	authMiddleware gin.HandlerFunc,
) *DriverRoutes {
	return &DriverRoutes{
		getByCityController: getByCity,
		authMiddleware:   authMiddleware,
	}
}

func (r *DriverRoutes) AttachRoutes(router *gin.Engine) {
	driversGroup := router.Group("/drivers")
	driversGroup.Use(r.authMiddleware)

	driversGroup.GET("/city/:city", r.getByCityController.GetByCity)
}
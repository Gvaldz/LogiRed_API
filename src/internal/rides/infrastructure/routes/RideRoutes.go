package routes

import (
	"logired/src/internal/rides/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type RideRoutes struct {
	createRideController          *controllers.CreateRideController
	cancelRideController          *controllers.CancelRideController
	getRidesByClientController    *controllers.GetRidesByClientController
	getRideByIdController         *controllers.GetRideByIdController
	getRidesByDriverController    *controllers.GetRidesByDriverController
	authMiddleware                gin.HandlerFunc
}

func NewRideRoutes(
	create 						  *controllers.CreateRideController,
	cancel 						  *controllers.CancelRideController,
	getByClient 				  *controllers.GetRidesByClientController,
	getById 					  *controllers.GetRideByIdController,
	getByDriver 				  *controllers.GetRidesByDriverController,
	authMiddleware 				  gin.HandlerFunc,
) *RideRoutes {
	return &RideRoutes{
		createRideController:          create,
		cancelRideController:          cancel,
		getRidesByClientController:    getByClient,
		getRideByIdController:         getById,
		getRidesByDriverController:    getByDriver,
		authMiddleware:                authMiddleware,
	}
}

func (r *RideRoutes) AttachRoutes(router *gin.Engine) {
	ridesGroup := router.Group("/rides")
	ridesGroup.Use(r.authMiddleware)

	ridesGroup.POST("", r.createRideController.Create)
	ridesGroup.DELETE("/:id", r.cancelRideController.Cancel)
	ridesGroup.GET("/client/me", r.getRidesByClientController.GetByClient)
	ridesGroup.GET("/:id", r.getRideByIdController.GetById)
	ridesGroup.GET("/driver/me", r.getRidesByDriverController.GetByDriver)
}
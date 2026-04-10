package routes

import (
	"logired/src/internal/rides/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type RideRoutes struct {
	createRideController          *controllers.CreateRideController
	getRidesByClientController    *controllers.GetRidesByClientController
	getRideByIdController         *controllers.GetRideByIdController
	getRidesByDriverController    *controllers.GetRidesByDriverController
	getRidesByCityController  	  *controllers.GetRideByCityController
	getRidesHistoryController     *controllers.GetRidesHistoryController
	updateRideStatusController    *controllers.UpdateRideStatusController
	authMiddleware                gin.HandlerFunc
}

func NewRideRoutes(
	create 						  *controllers.CreateRideController,
	getByClient 				  *controllers.GetRidesByClientController,
	getById 					  *controllers.GetRideByIdController,
	getByDriver 				  *controllers.GetRidesByDriverController,
	getByCity 					  *controllers.GetRideByCityController,
	getRidesHistory		  		  *controllers.GetRidesHistoryController,
	updateRideStatus		      *controllers.UpdateRideStatusController,
	authMiddleware 				  gin.HandlerFunc,
) *RideRoutes {
	return &RideRoutes{
		createRideController:          create,
		getRidesByClientController:    getByClient,
		getRideByIdController:         getById,
		getRidesByDriverController:    getByDriver,
		getRidesByCityController:  	   getByCity,
		getRidesHistoryController:     getRidesHistory,
		updateRideStatusController:    updateRideStatus,
		authMiddleware:                authMiddleware,
	}
}

func (r *RideRoutes) AttachRoutes(router *gin.Engine) {
	ridesGroup := router.Group("/rides")
	ridesGroup.Use(r.authMiddleware)

	ridesGroup.POST("", r.createRideController.Create)
	ridesGroup.GET("/client/me", r.getRidesByClientController.GetByClient)
	ridesGroup.GET("/:id", r.getRideByIdController.GetById)
	ridesGroup.GET("/driver/me", r.getRidesByDriverController.GetByDriver)
	ridesGroup.GET("/history", r.getRidesHistoryController.GetHistory)
	ridesGroup.GET("/city/:city", r.getRidesByCityController.GetByCity) 
	ridesGroup.PUT("/:id/status", r.updateRideStatusController.Update)
}
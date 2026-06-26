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
	getRidesByCityController      *controllers.GetRideByCityController
	getRidesHistoryController     *controllers.GetRidesHistoryController
	getRidesByStatusController    *controllers.GetRidesByStatusController
	updateRideStatusController    *controllers.UpdateRideStatusController
	authMiddleware                gin.HandlerFunc
	driverOnlyApprovedMiddleware  gin.HandlerFunc
}

func NewRideRoutes(
	create *controllers.CreateRideController,
	getByClient *controllers.GetRidesByClientController,
	getById *controllers.GetRideByIdController,
	getByDriver *controllers.GetRidesByDriverController,
	getByCity *controllers.GetRideByCityController,
	getRidesHistory *controllers.GetRidesHistoryController,
	getRidesByStatus *controllers.GetRidesByStatusController,
	updateRideStatus *controllers.UpdateRideStatusController,
	authMiddleware gin.HandlerFunc,
	driverOnlyApprovedMiddleware gin.HandlerFunc,
) *RideRoutes {
	return &RideRoutes{
		createRideController:         create,
		getRidesByClientController:   getByClient,
		getRideByIdController:        getById,
		getRidesByDriverController:   getByDriver,
		getRidesByCityController:     getByCity,
		getRidesHistoryController:    getRidesHistory,
		getRidesByStatusController:   getRidesByStatus,
		updateRideStatusController:   updateRideStatus,
		authMiddleware:               authMiddleware,
		driverOnlyApprovedMiddleware: driverOnlyApprovedMiddleware,
	}
}

func (r *RideRoutes) AttachRoutes(router *gin.Engine) {
	// Rutas públicas (requieren autenticación pero no aprobación de conductor)
	publicGroup := router.Group("/rides")
	publicGroup.Use(r.authMiddleware)
	publicGroup.POST("", r.createRideController.Create)  // Clientes crean viajes

	// Rutas SOLO para conductores aprobados
	driverGroup := router.Group("/rides")
	driverGroup.Use(r.authMiddleware)
	driverGroup.Use(r.driverOnlyApprovedMiddleware)
	driverGroup.GET("/available", r.getRidesByCityController.GetByCity)
	driverGroup.GET("/driver/me", r.getRidesByDriverController.GetByDriver)
	driverGroup.GET("/history", r.getRidesHistoryController.GetHistory)
	driverGroup.GET("/status", r.getRidesByStatusController.GetByStatus)

	// Rutas que ambos pueden usar (cliente + conductor)
	commonGroup := router.Group("/rides")
	commonGroup.Use(r.authMiddleware)
	commonGroup.GET("/:id", r.getRideByIdController.GetById)
	commonGroup.PUT("/:id/status", r.updateRideStatusController.Update)
	commonGroup.GET("/client/me", r.getRidesByClientController.GetByClient)
}
